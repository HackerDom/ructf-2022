from models import Contract, Procedure


class CreateDoctorReq:

    def __init__(self, doc_id, name, procedure_description, education_level):
        self.id = doc_id
        self.name = name
        self.procedure_description = procedure_description
        self.education_level = education_level

    def to_json(self):
        return {
            "Id": {"id": f"{self.id}"},
            "name": self.name,
            "ProcedureDescription": self.procedure_description,
            "EducationLevel": self.education_level
        }


class GetDoctorReq:
    def __init__(self, doc_id):
        self.id = doc_id

    def to_json(self):
        return {"id": f"{self.id}"}


class CreatePatientReq:
    def __init__(self, patient_id, name, diagnosis):
        self.id = patient_id
        self.name = name
        self.education_level = diagnosis

    def to_json(self):
        return {
            "Id": {"id": f"{self.id}"},
            "name": self.name,
            "diagnosis": self.education_level
        }


class CreateContractReq:
    def __init__(self, contract_id, doctor_id, doctor_signature, expired):
        self.id = contract_id
        self.doctor_id = doctor_id
        self.doctor_signature = doctor_signature
        self.expired = expired

    def to_json(self):
        return {
            "Id": {"id": f"{self.id}"},
            "DoctorId": {"id": f"{self.doctor_id}"},
            "DoctorSignature": {"Value": self.doctor_signature},
            "Expired": self.expired}


class PrescribeProcedureReq:
    def __init__(self, procedure_id, contract: Contract, procedure_type):
        self.id = procedure_id
        self.contract = contract
        self.procedure_type = procedure_type

    def to_json(self):
        return {
            "ContractModel": {
                "ContractId": {"id": f"{self.contract.contract_id}"},
                "Info": {
                    "PatientId": {"id": f"{self.contract.patient.patient_id}"},
                    "DoctorId": {"id": f"{self.contract.doctor.doc_id}"}
                },
                "Expired": f"{self.contract.expired}"
            },
            "ProcedureId": {"id": f"{self.id}"},
            "ProcedureType": self.procedure_type}


class PerformProcedureReq:
    def __init__(self, procedure_id: str):
        self.procedure_id = procedure_id

    def to_json(self):
        return {"id": f"{self.procedure_id}"}
