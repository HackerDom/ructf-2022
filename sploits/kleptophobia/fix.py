import models.models_pb2 as m


def fix(pp_str, k):
    return b'\x12' + bytes([k]) + b'A'*k + pp_str


if __name__ == '__main__':
    pp = m.PrivatePerson(
        first_name = "first_name",
        middle_name = "middle_name",
        second_name = "second_name",
        room = 12345,
        diagnosis = "diagnosis"
    )

    pp_str = pp.SerializeToString()
    pp_str = fix(pp_str, 5)

    fixed_pp = m.PrivatePerson()
    fixed_pp.ParseFromString(pp_str)
    for field in ['first_name', 'middle_name', 'second_name', 'room', 'diagnosis']:
        assert getattr(pp, field) == getattr(fixed_pp, field)
    print('fixed')