# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/screen.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/screen.proto',
  package='laptop',
  syntax='proto3',
  serialized_options=b'Z\013.;___laptop',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x12proto/screen.proto\x12\x06laptop\"\xd4\x01\n\x06Screen\x12\x0c\n\x04size\x18\x01 \x01(\x02\x12-\n\nresolution\x18\x02 \x01(\x0b\x32\x19.laptop.Screen.Resolution\x12#\n\x05panel\x18\x03 \x01(\x0e\x32\x14.laptop.Screen.Panel\x12\x12\n\nmultitouch\x18\x04 \x01(\x08\x1a+\n\nResolution\x12\r\n\x05width\x18\x01 \x01(\r\x12\x0e\n\x06height\x18\x02 \x01(\r\"\'\n\x05Panel\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x07\n\x03IPS\x10\x01\x12\x08\n\x04OLED\x10\x02\x42\rZ\x0b.;___laptopb\x06proto3'
)



_SCREEN_PANEL = _descriptor.EnumDescriptor(
  name='Panel',
  full_name='laptop.Screen.Panel',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='IPS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OLED', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=204,
  serialized_end=243,
)
_sym_db.RegisterEnumDescriptor(_SCREEN_PANEL)


_SCREEN_RESOLUTION = _descriptor.Descriptor(
  name='Resolution',
  full_name='laptop.Screen.Resolution',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='width', full_name='laptop.Screen.Resolution.width', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='height', full_name='laptop.Screen.Resolution.height', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=159,
  serialized_end=202,
)

_SCREEN = _descriptor.Descriptor(
  name='Screen',
  full_name='laptop.Screen',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='size', full_name='laptop.Screen.size', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resolution', full_name='laptop.Screen.resolution', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='panel', full_name='laptop.Screen.panel', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='multitouch', full_name='laptop.Screen.multitouch', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_SCREEN_RESOLUTION, ],
  enum_types=[
    _SCREEN_PANEL,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=243,
)

_SCREEN_RESOLUTION.containing_type = _SCREEN
_SCREEN.fields_by_name['resolution'].message_type = _SCREEN_RESOLUTION
_SCREEN.fields_by_name['panel'].enum_type = _SCREEN_PANEL
_SCREEN_PANEL.containing_type = _SCREEN
DESCRIPTOR.message_types_by_name['Screen'] = _SCREEN
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Screen = _reflection.GeneratedProtocolMessageType('Screen', (_message.Message,), {

  'Resolution' : _reflection.GeneratedProtocolMessageType('Resolution', (_message.Message,), {
    'DESCRIPTOR' : _SCREEN_RESOLUTION,
    '__module__' : 'proto.screen_pb2'
    # @@protoc_insertion_point(class_scope:laptop.Screen.Resolution)
    })
  ,
  'DESCRIPTOR' : _SCREEN,
  '__module__' : 'proto.screen_pb2'
  # @@protoc_insertion_point(class_scope:laptop.Screen)
  })
_sym_db.RegisterMessage(Screen)
_sym_db.RegisterMessage(Screen.Resolution)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)