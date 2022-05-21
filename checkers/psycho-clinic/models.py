import random
from uuid import uuid4

from utils import get_rnd_string, get_rnd_future_date

edu_levels = ["Master", "Bachelor", "PhD"]
diagnosis_types = ["Schizophrenia", "PanicDisorder", "ObsessiveCompulsiveDisorder", "Depression", "BipolarDisorder"]
procedure_types = ["MedicationInjection", "Psychotherapy", "Psychoanalysis"]


class Doctor:
    def __init__(self, doc_id: str, name: str, proc_desc: str, edu_lvl: str, signature: str):
        self.doc_id = doc_id
        self.name = name
        self.proc_desc = proc_desc
        self.edu_lvl = edu_lvl
        self.signature = signature

    @staticmethod
    def gen(with_sign=False):
        return Doctor(
            str(uuid4()),
            get_rnd_string(15),
            get_rnd_string(50),
            random.choice(edu_levels),
            get_rnd_string(40) if with_sign else None)

    def compare(self, doc_json, with_sign=True):
        return self.doc_id == doc_json["id"]["id"] and \
               self.name == doc_json["name"] and \
               self.edu_lvl == doc_json["educationLevel"] and \
               self.proc_desc == doc_json["procedureDescription"] and \
               self.signature == doc_json["signature"]["value"] if with_sign else True


class Patient:
    def __init__(self, patient_id: str, name, diagnosis: str, patient_token: str):
        self.patient_id = patient_id
        self.name = name
        self.diagnosis = diagnosis
        self.patient_token = patient_token

    @staticmethod
    def gen(with_sign=False):
        return Patient(
            str(uuid4()),
            get_rnd_string(15),
            random.choice(diagnosis_types),
            get_rnd_string(40) if with_sign else None)


class Contract:
    def __init__(self, contract_id: str, patient: Patient, doctor: Doctor, expired: str):
        self.contract_id = contract_id
        self.patient = patient
        self.doctor = doctor
        self.expired = expired

    @staticmethod
    def gen(patient: Patient, doctor: Doctor):
        return Contract(
            str(uuid4()),
            patient,
            doctor,
            str(get_rnd_future_date()))


class Procedure:
    def __init__(self, procedure_id: str, contract: Contract, procedure_type: str):
        self.procedure_id = procedure_id
        self.contract = contract
        self.procedure_type = procedure_type

    @staticmethod
    def gen(contract: Contract):
        return Procedure(
            str(uuid4()),
            contract,
            random.choice(procedure_types))


class ProcedurePerformResult:
    def __init__(self, procedure_id: str, is_successful: bool, description: str):
        self.procedure_id = procedure_id
        self.is_successful = is_successful
        self.description = description
