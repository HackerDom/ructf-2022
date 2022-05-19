# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cmodels.proto\x12\x06models\"\x1b\n\x08PingBody\x12\x0f\n\x07message\x18\x01 \x01(\t\"$\n\x10GetByUsernameReq\x12\x10\n\x08username\x18\x01 \x01(\t\"\xb7\x01\n\x10GetPublicInfoRsp\x12/\n\x06status\x18\x01 \x01(\x0e\x32\x1f.models.GetPublicInfoRsp.Status\x12\x14\n\x07message\x18\x02 \x01(\tH\x00\x88\x01\x01\x12)\n\x06person\x18\x03 \x01(\x0b\x32\x14.models.PublicPersonH\x01\x88\x01\x01\"\x1a\n\x06Status\x12\x06\n\x02OK\x10\x00\x12\x08\n\x04\x46\x41IL\x10\x01\x42\n\n\x08_messageB\t\n\x07_person\"\xc5\x01\n\x17GetEncryptedFullInfoRsp\x12\x36\n\x06status\x18\x01 \x01(\x0e\x32&.models.GetEncryptedFullInfoRsp.Status\x12\x14\n\x07message\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x1e\n\x11\x65ncryptedFullInfo\x18\x03 \x01(\x0cH\x01\x88\x01\x01\"\x1a\n\x06Status\x12\x06\n\x02OK\x10\x00\x12\x08\n\x04\x46\x41IL\x10\x01\x42\n\n\x08_messageB\x14\n\x12_encryptedFullInfo\"F\n\x0bRegisterReq\x12%\n\x06person\x18\x01 \x01(\x0b\x32\x15.models.PrivatePerson\x12\x10\n\x08password\x18\x02 \x01(\t\"w\n\x0bRegisterRsp\x12*\n\x06status\x18\x01 \x01(\x0e\x32\x1a.models.RegisterRsp.Status\x12\x14\n\x07message\x18\x02 \x01(\tH\x00\x88\x01\x01\"\x1a\n\x06Status\x12\x06\n\x02OK\x10\x00\x12\x08\n\x04\x46\x41IL\x10\x01\x42\n\n\x08_message\"\x95\x01\n\rPrivatePerson\x12\x12\n\nfirst_name\x18\x01 \x01(\t\x12\x18\n\x0bmiddle_name\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x13\n\x0bsecond_name\x18\x03 \x01(\t\x12\x10\n\x08username\x18\x04 \x01(\t\x12\x0c\n\x04room\x18\x05 \x01(\x05\x12\x11\n\tdiagnosis\x18\x06 \x01(\tB\x0e\n\x0c_middle_name\"W\n\x0cPublicPerson\x12\x12\n\nfirst_name\x18\x01 \x01(\t\x12\x13\n\x0bsecond_name\x18\x02 \x01(\t\x12\x10\n\x08username\x18\x03 \x01(\t\x12\x0c\n\x04room\x18\x04 \x01(\x05\"4\n\x0c\x43lientConfig\x12\x11\n\tgrpc_host\x18\x01 \x01(\t\x12\x11\n\tgrpc_port\x18\x02 \x01(\x05\"[\n\x08PGConfig\x12\x0c\n\x04host\x18\x01 \x01(\t\x12\x0c\n\x04port\x18\x02 \x01(\x05\x12\x10\n\x08username\x18\x03 \x01(\t\x12\x10\n\x08password\x18\x04 \x01(\t\x12\x0f\n\x07\x64\x62_name\x18\x05 \x01(\t\"F\n\x0cServerConfig\x12\x11\n\tgrpc_port\x18\x01 \x01(\x05\x12#\n\tpg_config\x18\x02 \x01(\x0b\x32\x10.models.PGConfig2\x90\x02\n\x0cKleptophobia\x12\x36\n\x08Register\x12\x13.models.RegisterReq\x1a\x13.models.RegisterRsp\"\x00\x12\x45\n\rGetPublicInfo\x12\x18.models.GetByUsernameReq\x1a\x18.models.GetPublicInfoRsp\"\x00\x12S\n\x14GetEncryptedFullInfo\x12\x18.models.GetByUsernameReq\x1a\x1f.models.GetEncryptedFullInfoRsp\"\x00\x12,\n\x04Ping\x12\x10.models.PingBody\x1a\x10.models.PingBody\"\x00\x42\x16Z\x14hackerdom.com/modelsb\x06proto3')



_PINGBODY = DESCRIPTOR.message_types_by_name['PingBody']
_GETBYUSERNAMEREQ = DESCRIPTOR.message_types_by_name['GetByUsernameReq']
_GETPUBLICINFORSP = DESCRIPTOR.message_types_by_name['GetPublicInfoRsp']
_GETENCRYPTEDFULLINFORSP = DESCRIPTOR.message_types_by_name['GetEncryptedFullInfoRsp']
_REGISTERREQ = DESCRIPTOR.message_types_by_name['RegisterReq']
_REGISTERRSP = DESCRIPTOR.message_types_by_name['RegisterRsp']
_PRIVATEPERSON = DESCRIPTOR.message_types_by_name['PrivatePerson']
_PUBLICPERSON = DESCRIPTOR.message_types_by_name['PublicPerson']
_CLIENTCONFIG = DESCRIPTOR.message_types_by_name['ClientConfig']
_PGCONFIG = DESCRIPTOR.message_types_by_name['PGConfig']
_SERVERCONFIG = DESCRIPTOR.message_types_by_name['ServerConfig']
_GETPUBLICINFORSP_STATUS = _GETPUBLICINFORSP.enum_types_by_name['Status']
_GETENCRYPTEDFULLINFORSP_STATUS = _GETENCRYPTEDFULLINFORSP.enum_types_by_name['Status']
_REGISTERRSP_STATUS = _REGISTERRSP.enum_types_by_name['Status']
PingBody = _reflection.GeneratedProtocolMessageType('PingBody', (_message.Message,), {
  'DESCRIPTOR' : _PINGBODY,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.PingBody)
  })
_sym_db.RegisterMessage(PingBody)

GetByUsernameReq = _reflection.GeneratedProtocolMessageType('GetByUsernameReq', (_message.Message,), {
  'DESCRIPTOR' : _GETBYUSERNAMEREQ,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.GetByUsernameReq)
  })
_sym_db.RegisterMessage(GetByUsernameReq)

GetPublicInfoRsp = _reflection.GeneratedProtocolMessageType('GetPublicInfoRsp', (_message.Message,), {
  'DESCRIPTOR' : _GETPUBLICINFORSP,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.GetPublicInfoRsp)
  })
_sym_db.RegisterMessage(GetPublicInfoRsp)

GetEncryptedFullInfoRsp = _reflection.GeneratedProtocolMessageType('GetEncryptedFullInfoRsp', (_message.Message,), {
  'DESCRIPTOR' : _GETENCRYPTEDFULLINFORSP,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.GetEncryptedFullInfoRsp)
  })
_sym_db.RegisterMessage(GetEncryptedFullInfoRsp)

RegisterReq = _reflection.GeneratedProtocolMessageType('RegisterReq', (_message.Message,), {
  'DESCRIPTOR' : _REGISTERREQ,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.RegisterReq)
  })
_sym_db.RegisterMessage(RegisterReq)

RegisterRsp = _reflection.GeneratedProtocolMessageType('RegisterRsp', (_message.Message,), {
  'DESCRIPTOR' : _REGISTERRSP,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.RegisterRsp)
  })
_sym_db.RegisterMessage(RegisterRsp)

PrivatePerson = _reflection.GeneratedProtocolMessageType('PrivatePerson', (_message.Message,), {
  'DESCRIPTOR' : _PRIVATEPERSON,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.PrivatePerson)
  })
_sym_db.RegisterMessage(PrivatePerson)

PublicPerson = _reflection.GeneratedProtocolMessageType('PublicPerson', (_message.Message,), {
  'DESCRIPTOR' : _PUBLICPERSON,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.PublicPerson)
  })
_sym_db.RegisterMessage(PublicPerson)

ClientConfig = _reflection.GeneratedProtocolMessageType('ClientConfig', (_message.Message,), {
  'DESCRIPTOR' : _CLIENTCONFIG,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.ClientConfig)
  })
_sym_db.RegisterMessage(ClientConfig)

PGConfig = _reflection.GeneratedProtocolMessageType('PGConfig', (_message.Message,), {
  'DESCRIPTOR' : _PGCONFIG,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.PGConfig)
  })
_sym_db.RegisterMessage(PGConfig)

ServerConfig = _reflection.GeneratedProtocolMessageType('ServerConfig', (_message.Message,), {
  'DESCRIPTOR' : _SERVERCONFIG,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:models.ServerConfig)
  })
_sym_db.RegisterMessage(ServerConfig)

_KLEPTOPHOBIA = DESCRIPTOR.services_by_name['Kleptophobia']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\024hackerdom.com/models'
  _PINGBODY._serialized_start=24
  _PINGBODY._serialized_end=51
  _GETBYUSERNAMEREQ._serialized_start=53
  _GETBYUSERNAMEREQ._serialized_end=89
  _GETPUBLICINFORSP._serialized_start=92
  _GETPUBLICINFORSP._serialized_end=275
  _GETPUBLICINFORSP_STATUS._serialized_start=226
  _GETPUBLICINFORSP_STATUS._serialized_end=252
  _GETENCRYPTEDFULLINFORSP._serialized_start=278
  _GETENCRYPTEDFULLINFORSP._serialized_end=475
  _GETENCRYPTEDFULLINFORSP_STATUS._serialized_start=226
  _GETENCRYPTEDFULLINFORSP_STATUS._serialized_end=252
  _REGISTERREQ._serialized_start=477
  _REGISTERREQ._serialized_end=547
  _REGISTERRSP._serialized_start=549
  _REGISTERRSP._serialized_end=668
  _REGISTERRSP_STATUS._serialized_start=226
  _REGISTERRSP_STATUS._serialized_end=252
  _PRIVATEPERSON._serialized_start=671
  _PRIVATEPERSON._serialized_end=820
  _PUBLICPERSON._serialized_start=822
  _PUBLICPERSON._serialized_end=909
  _CLIENTCONFIG._serialized_start=911
  _CLIENTCONFIG._serialized_end=963
  _PGCONFIG._serialized_start=965
  _PGCONFIG._serialized_end=1056
  _SERVERCONFIG._serialized_start=1058
  _SERVERCONFIG._serialized_end=1128
  _KLEPTOPHOBIA._serialized_start=1131
  _KLEPTOPHOBIA._serialized_end=1403
# @@protoc_insertion_point(module_scope)