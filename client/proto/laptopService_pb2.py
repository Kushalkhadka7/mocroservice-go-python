# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/laptopService.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto import laptop_pb2 as proto_dot_laptop__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/laptopService.proto',
  package='laptop',
  syntax='proto3',
  serialized_options=b'Z\013.;___laptop',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x19proto/laptopService.proto\x12\x06laptop\x1a\x12proto/laptop.proto\"5\n\x13\x43reateLaptopRequest\x12\x1e\n\x06laptop\x18\x01 \x01(\x0b\x32\x0e.laptop.Laptop\"6\n\x14\x43reateLaptopResponse\x12\x1e\n\x06laptop\x18\x01 \x01(\x0b\x32\x0e.laptop.Laptop\"\x16\n\x05Hello\x12\r\n\x05Hello\x18\x01 \x01(\t2\x97\x01\n\rLaptopService\x12K\n\x0c\x43reateLaptop\x12\x1b.laptop.CreateLaptopRequest\x1a\x1c.laptop.CreateLaptopResponse\"\x00\x12\x39\n\x08SayHello\x12\r.laptop.Hello\x1a\x1c.laptop.CreateLaptopResponse\"\x00\x42\rZ\x0b.;___laptopb\x06proto3'
  ,
  dependencies=[proto_dot_laptop__pb2.DESCRIPTOR,])




_CREATELAPTOPREQUEST = _descriptor.Descriptor(
  name='CreateLaptopRequest',
  full_name='laptop.CreateLaptopRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='laptop', full_name='laptop.CreateLaptopRequest.laptop', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=110,
)


_CREATELAPTOPRESPONSE = _descriptor.Descriptor(
  name='CreateLaptopResponse',
  full_name='laptop.CreateLaptopResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='laptop', full_name='laptop.CreateLaptopResponse.laptop', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=112,
  serialized_end=166,
)


_HELLO = _descriptor.Descriptor(
  name='Hello',
  full_name='laptop.Hello',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Hello', full_name='laptop.Hello.Hello', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=168,
  serialized_end=190,
)

_CREATELAPTOPREQUEST.fields_by_name['laptop'].message_type = proto_dot_laptop__pb2._LAPTOP
_CREATELAPTOPRESPONSE.fields_by_name['laptop'].message_type = proto_dot_laptop__pb2._LAPTOP
DESCRIPTOR.message_types_by_name['CreateLaptopRequest'] = _CREATELAPTOPREQUEST
DESCRIPTOR.message_types_by_name['CreateLaptopResponse'] = _CREATELAPTOPRESPONSE
DESCRIPTOR.message_types_by_name['Hello'] = _HELLO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateLaptopRequest = _reflection.GeneratedProtocolMessageType('CreateLaptopRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATELAPTOPREQUEST,
  '__module__' : 'proto.laptopService_pb2'
  # @@protoc_insertion_point(class_scope:laptop.CreateLaptopRequest)
  })
_sym_db.RegisterMessage(CreateLaptopRequest)

CreateLaptopResponse = _reflection.GeneratedProtocolMessageType('CreateLaptopResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATELAPTOPRESPONSE,
  '__module__' : 'proto.laptopService_pb2'
  # @@protoc_insertion_point(class_scope:laptop.CreateLaptopResponse)
  })
_sym_db.RegisterMessage(CreateLaptopResponse)

Hello = _reflection.GeneratedProtocolMessageType('Hello', (_message.Message,), {
  'DESCRIPTOR' : _HELLO,
  '__module__' : 'proto.laptopService_pb2'
  # @@protoc_insertion_point(class_scope:laptop.Hello)
  })
_sym_db.RegisterMessage(Hello)


DESCRIPTOR._options = None

_LAPTOPSERVICE = _descriptor.ServiceDescriptor(
  name='LaptopService',
  full_name='laptop.LaptopService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=193,
  serialized_end=344,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateLaptop',
    full_name='laptop.LaptopService.CreateLaptop',
    index=0,
    containing_service=None,
    input_type=_CREATELAPTOPREQUEST,
    output_type=_CREATELAPTOPRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SayHello',
    full_name='laptop.LaptopService.SayHello',
    index=1,
    containing_service=None,
    input_type=_HELLO,
    output_type=_CREATELAPTOPRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_LAPTOPSERVICE)

DESCRIPTOR.services_by_name['LaptopService'] = _LAPTOPSERVICE

# @@protoc_insertion_point(module_scope)
