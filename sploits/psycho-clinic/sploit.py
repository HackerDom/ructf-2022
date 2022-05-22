from uuid import uuid4

from client import Client
from models import Contract
from utils import get_rnd_future_date


def main():
    client = Client("localhost")

    FLAG = f"{uuid4()}"
    print(f"Flag is: {FLAG}")

    patient_name = "some_name"
    doctor_name = "doc_name"
    secret_doctor_name = "secret_doc_name"
    doctor_desc = "some_desc"
    contract_expired = str(get_rnd_future_date())
    some_edu_lvl = "Master"
    some_procedure_type = "MedicationInjection"
    some_diagnosis = "Schizophrenia"

    patient = client.create_patient(str(uuid4()), patient_name, some_diagnosis)

    doctor_secret = client.create_doctor(str(uuid4()), secret_doctor_name, FLAG, some_edu_lvl)
    doctor_common = client.create_doctor(str(uuid4()), doctor_name, doctor_desc, some_edu_lvl)

    fake_contract = client.create_contract(str(uuid4()), patient, doctor_common, contract_expired)
    hacked_contract = Contract(fake_contract.contract_id, patient, doctor_secret, contract_expired)

    hacked_procedure = client.prescribe_procedure(str(uuid4()), hacked_contract, some_procedure_type)
    hacked_procedure_perform_result = client.perform_procedure(hacked_procedure.procedure_id, patient.patient_token)

    potential_flag = hacked_procedure_perform_result.description
    if potential_flag == FLAG:
        print(f"CONGRATS! Your reached the flag: {potential_flag}")
    else:
        raise Exception("smth went wrong :(")


def iterative_main():
    for i in range(0, 5):
        main()
        print()


if __name__ == '__main__':
    iterative_main()
