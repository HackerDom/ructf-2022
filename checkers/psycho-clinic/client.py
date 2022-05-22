import traceback

from gornilo import Verdict
from requests import post, get

from models import Patient, Doctor, ProcedurePerformResult
from request_models import *
from utils import raise_not_found_exc, raise_data_exc, VerdictHttpException


class Client:
    def __init__(self, host, port=18323):
        self.app_prefix = f"http://{host}:{port}/"

    # region Requests sending
    def _send_request(self, url, request, json_data=None, api_key=None, params=None):
        headers = {"Authorization": f"Basic {api_key}"} if api_key is not None else None

        absolute_url = self.app_prefix + url

        try:
            response = request(url=absolute_url, json=json_data, headers=headers, params=params)
        except Exception as e:
            raise VerdictHttpException(Verdict.DOWN("HTTP error."))

        if response.status_code != 200:
            message = f"Invalid status code: {response.status_code} for {url}."
            private_msg = f"resp: {response.text}. was send: {json_data}"
            raise VerdictHttpException(Verdict.MUMBLE(message), private_msg)

        return response

    def send_create_doctor(self, doc_id: str, name: str, proc_d: str, edu_level: str):
        url = "doctors/create"
        return self._send_request(url, post, json_data=CreateDoctorReq(doc_id, name, proc_d, edu_level).to_json())

    def send_get_doctor(self, doc_id: str):
        url = "doctors/"
        return self._send_request(url, post, json_data=GetDoctorReq(doc_id).to_json()).json()

    def send_get_doctors(self, edu_lvl: str, skip: int = 0, take: int = 10):
        url = "doctors/"
        params = {"educationLevel": edu_lvl, "skip": str(skip), "take": str(take)}
        return self._send_request(url, get, params=params).json()

    def send_create_patient(self, patient_id: str, name: str, diagnosis: str):
        url = "patient/create"
        return self._send_request(url, post, json_data=CreatePatientReq(patient_id, name, diagnosis).to_json())

    def send_get_patient(self, patient_token: str):
        url = "patient/card"
        return self._send_request(url, get, api_key=patient_token).json()

    def send_create_contract(self, contract_id: str, patient: Patient, doctor: Doctor, expired: str):
        url = "contracts/create"

        req = CreateContractReq(contract_id, doctor.doc_id, doctor.signature, expired)

        return self._send_request(url, post, json_data=req.to_json(), api_key=patient.patient_token)

    def send_get_contracts(self, patient: Patient) -> list:
        url = "contracts/all"
        return self._send_request(url, get, api_key=patient.patient_token).json()

    def send_prescribe_procedure(self, procedure_id: str, contract: Contract, procedure_type: str):
        url = "procedures/prescribe"

        req = PrescribeProcedureReq(procedure_id, contract, procedure_type)

        return self._send_request(url, post, json_data=req.to_json(), api_key=contract.patient.patient_token)

    def send_get_procedures(self, patient: Patient):
        url = "procedures/all"
        return self._send_request(url, get, api_key=patient.patient_token).json()

    def send_perform_procedure(self, procedure_id: str, patient_token: str):
        url = "procedures/perform"

        resp = self._send_request(url,
                                  post,
                                  json_data=PerformProcedureReq(procedure_id).to_json(),
                                  api_key=patient_token)

        return resp.json()

    # endregion

    def create_doctor(self, doc_id: str, name: str, proc_desc: str, edu_lvl: str) -> Doctor:
        resp = self.send_create_doctor(doc_id, name, proc_desc, edu_lvl).json()

        doc_id = resp["id"]["id"]
        doc_signature = resp["signature"]["value"]
        result = Doctor(doc_id, name, proc_desc, edu_lvl, doc_signature)

        received_doc = self.send_get_doctor(doc_id)

        if not result.compare(received_doc, False):
            raise_data_exc("doctor", received_doc)

        return result

    def create_patient(self, patient_id: str, name: str, diagnosis: str) -> Patient:
        resp = self.send_create_patient(patient_id, name, diagnosis)
        patient_token = resp.json()["value"]

        self.send_get_patient(patient_token)

        return Patient(patient_id, name, diagnosis, patient_token)

    def create_contract(self, contract_id: str, patient: Patient, doctor: Doctor, expired: str) -> Contract:
        self.send_create_contract(contract_id, patient, doctor, expired)

        contracts = self.send_get_contracts(patient)
        ids = [x["contractId"]["id"] for x in contracts]

        if contract_id not in ids:
            raise_not_found_exc("Contract", contract_id)

        return Contract(contract_id, patient, doctor, expired)

    def prescribe_procedure(self, procedure_id: str, contract: Contract, procedure_type: str) -> Procedure:
        self.send_prescribe_procedure(procedure_id, contract, procedure_type)

        procedures = self.send_get_procedures(contract.patient)
        ids = [x["id"]["id"] for x in procedures]

        if f"{procedure_id}" not in ids:
            raise_not_found_exc("Procedure", procedure_id)

        return Procedure(procedure_id, contract, procedure_type)

    def perform_procedure(self, procedure_id: str, patient_token: str) -> ProcedurePerformResult:
        resp = self.send_perform_procedure(procedure_id, patient_token)

        is_successful = resp["procedureResult"]["isSuccessful"]
        description = resp["procedureResult"]["description"]

        return ProcedurePerformResult(procedure_id, is_successful, description)
