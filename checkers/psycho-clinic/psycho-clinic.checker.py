#!/usr/bin/env python3

import json
import time
import traceback

from gornilo import Checker, Verdict, PutRequest, GetRequest, CheckRequest

from client import Client
from models import Patient, Doctor, Contract, Procedure
from utils import raise_not_found_exc, VerdictDataException, VerdictHttpException, VerdictNotFoundException

checker = Checker()


@checker.define_check
def check_service(request: CheckRequest) -> Verdict:
    client = Client(request.hostname)

    patient = Patient.gen()
    doc1 = Doctor.gen()

    try:
        print("checking Patient")
        patient = client.create_patient(patient.patient_id, patient.name, patient.diagnosis)

        print("checking Doctor")
        doc1 = client.create_doctor(doc1.doc_id, doc1.name, doc1.proc_desc, doc1.edu_lvl)
        check_get_doctors(client, doc1)

        print("checking Contract")
        contract = Contract.gen(patient, doc1)
        client.create_contract(contract.contract_id, contract.patient, contract.doctor, contract.expired)

        print("checking Procedures")
        procedure = Procedure.gen(contract)
        client.prescribe_procedure(procedure.procedure_id, contract, procedure.procedure_type)
        client.perform_procedure(procedure.procedure_id, patient.patient_token)

        return Verdict.OK()
    except VerdictDataException as e:
        print(e)
        return e.verdict
    except VerdictHttpException as e:
        print(e)
        return e.verdict
    except VerdictNotFoundException as e:
        print(e)
        return e.verdict
    except Exception as e:
        traceback.print_exc()
        return Verdict.CORRUPT("Corrupted")


@checker.define_put(vuln_num=1, vuln_rate=1)
def put_flag(request: PutRequest) -> Verdict:
    client = Client(request.hostname)
    flag = request.flag

    patient = Patient.gen()
    doc_with_flag = Doctor.gen()

    try:
        print(f"Create doctor {doc_with_flag.doc_id}{doc_with_flag.edu_lvl}")
        doc_with_flag = client.create_doctor(doc_with_flag.doc_id, doc_with_flag.name, flag, doc_with_flag.edu_lvl)

        print(f"Create patient {patient.patient_id}{patient.diagnosis}")
        patient = client.create_patient(patient.patient_id, patient.name, patient.diagnosis)

        contract = Contract.gen(patient, doc_with_flag)
        print(f"Create contract {contract.contract_id}")
        client.create_contract(contract.contract_id, contract.patient, contract.doctor, contract.expired)

        procedure = Procedure.gen(contract)
        print(f"Prescribe procedure {procedure.procedure_id}")
        client.prescribe_procedure(procedure.procedure_id, contract, procedure.procedure_type)

        print("Saved flag " + flag)
        flag_id = json.dumps({"patient_token": patient.patient_token,
                              "doctor_id": doc_with_flag.doc_id,
                              "procedure_id": procedure.procedure_id})

        print("Saved flag_id " + flag_id)

        return Verdict.OK(flag_id)
    except VerdictHttpException as e:
        print(e)
        return e.verdict
    except Exception as e:
        traceback.print_exc()
        return Verdict.MUMBLE("Couldn't put flag!")


@checker.define_get(vuln_num=1)
def get_flag(request: GetRequest) -> Verdict:
    client = Client(request.hostname)
    flag = request.flag
    flag_id = json.loads(request.flag_id)

    try:
        patient = Patient.gen()
        patient.patient_token = flag_id["patient_token"]

        doctor = Doctor.gen()
        doctor.doc_id = flag_id["doctor_id"]

        print(f'Performing procedure {flag_id["procedure_id"]} for {patient.patient_token}')
        result = client.perform_procedure(flag_id["procedure_id"], patient.patient_token)
        if result.description == flag:
            return Verdict.OK()

        print("Procedure description doesn't contain a correct flag.")
        print(f"expected: '{flag}' but was: '{result.description}'")

        return Verdict.MUMBLE("Flag is missing!")
    except VerdictHttpException as e:
        print(e)
        return e.verdict
    except Exception as e:
        traceback.print_exc()
        return Verdict.MUMBLE("Couldn't get flag!")


def check_get_doctors(client, doctor_to_find: Doctor):
    response = client.send_get_doctors(doctor_to_find.edu_lvl)

    count, doctors = response["count"], response["doctors"]
    print(f"{count} doctors was found")
    take = len(doctors)
    skip = 0

    while skip < count:
        ids = [x["id"]["id"] for x in doctors]
        if doctor_to_find.doc_id in ids:
            return

        time.sleep(1.5)

        skip += take
        doctors = client.send_get_doctors(doctor_to_find.edu_lvl)["doctors"]

    raise_not_found_exc("Doctor", doctor_to_find.doc_id)


if __name__ == '__main__':
    checker.run()
