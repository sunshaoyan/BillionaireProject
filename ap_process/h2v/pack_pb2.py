# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pack.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='pack.proto',
  package='pack',
  syntax='proto3',
  serialized_options=_b('H\003'),
  serialized_pb=_b('\n\npack.proto\x12\x04pack\"\xa5\x01\n\x05\x46rame\x12\x14\n\x0csequence_id_\x18\x01 \x01(\x04\x12\x12\n\ntimestamp_\x18\x02 \x01(\x04\x12*\n\x0b\x66rame_type_\x18\x03 \x01(\x0e\x32\x15.pack.Frame.FrameType\"F\n\tFrameType\x12\n\n\x06Unkown\x10\x00\x12\x0e\n\nSmartFrame\x10\x01\x12\r\n\tDropFrame\x10\x02\x12\x0e\n\nErrorFrame\x10\x03\"\x17\n\x05\x43heck\x12\x0e\n\x06md5sum\x18\x01 \x03(\x0c\"D\n\x08\x41\x64\x64ition\x12\x1b\n\x06\x63heck_\x18\x01 \x01(\x0b\x32\x0b.pack.Check\x12\x1b\n\x06\x66rame_\x18\x02 \x01(\x0b\x32\x0b.pack.Frame\"\xed\x01\n\x0bMessagePack\x12%\n\x05\x66low_\x18\x01 \x01(\x0e\x32\x16.pack.MessagePack.Flow\x12%\n\x05type_\x18\x02 \x01(\x0e\x32\x16.pack.MessagePack.Type\x12!\n\taddition_\x18\x03 \x01(\x0b\x32\x0e.pack.Addition\x12\x10\n\x08\x63ontent_\x18\x04 \x01(\x0c\")\n\x04\x46low\x12\x0b\n\x07Unknown\x10\x00\x12\t\n\x05\x41P2CP\x10\x01\x12\t\n\x05\x43P2AP\x10\x02\"0\n\x04Type\x12\x0c\n\x08kUnknown\x10\x00\x12\x0c\n\x08kXPlugin\x10\x01\x12\x0c\n\x08kXConfig\x10\x02\x42\x02H\x03\x62\x06proto3')
)



_FRAME_FRAMETYPE = _descriptor.EnumDescriptor(
  name='FrameType',
  full_name='pack.Frame.FrameType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Unkown', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SmartFrame', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DropFrame', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ErrorFrame', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=116,
  serialized_end=186,
)
_sym_db.RegisterEnumDescriptor(_FRAME_FRAMETYPE)

_MESSAGEPACK_FLOW = _descriptor.EnumDescriptor(
  name='Flow',
  full_name='pack.MessagePack.Flow',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Unknown', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP2CP', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CP2AP', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=430,
  serialized_end=471,
)
_sym_db.RegisterEnumDescriptor(_MESSAGEPACK_FLOW)

_MESSAGEPACK_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='pack.MessagePack.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='kUnknown', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='kXPlugin', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='kXConfig', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=473,
  serialized_end=521,
)
_sym_db.RegisterEnumDescriptor(_MESSAGEPACK_TYPE)


_FRAME = _descriptor.Descriptor(
  name='Frame',
  full_name='pack.Frame',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sequence_id_', full_name='pack.Frame.sequence_id_', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp_', full_name='pack.Frame.timestamp_', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='frame_type_', full_name='pack.Frame.frame_type_', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _FRAME_FRAMETYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=21,
  serialized_end=186,
)


_CHECK = _descriptor.Descriptor(
  name='Check',
  full_name='pack.Check',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='md5sum', full_name='pack.Check.md5sum', index=0,
      number=1, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=188,
  serialized_end=211,
)


_ADDITION = _descriptor.Descriptor(
  name='Addition',
  full_name='pack.Addition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='check_', full_name='pack.Addition.check_', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='frame_', full_name='pack.Addition.frame_', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=213,
  serialized_end=281,
)


_MESSAGEPACK = _descriptor.Descriptor(
  name='MessagePack',
  full_name='pack.MessagePack',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='flow_', full_name='pack.MessagePack.flow_', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type_', full_name='pack.MessagePack.type_', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addition_', full_name='pack.MessagePack.addition_', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='content_', full_name='pack.MessagePack.content_', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _MESSAGEPACK_FLOW,
    _MESSAGEPACK_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=284,
  serialized_end=521,
)

_FRAME.fields_by_name['frame_type_'].enum_type = _FRAME_FRAMETYPE
_FRAME_FRAMETYPE.containing_type = _FRAME
_ADDITION.fields_by_name['check_'].message_type = _CHECK
_ADDITION.fields_by_name['frame_'].message_type = _FRAME
_MESSAGEPACK.fields_by_name['flow_'].enum_type = _MESSAGEPACK_FLOW
_MESSAGEPACK.fields_by_name['type_'].enum_type = _MESSAGEPACK_TYPE
_MESSAGEPACK.fields_by_name['addition_'].message_type = _ADDITION
_MESSAGEPACK_FLOW.containing_type = _MESSAGEPACK
_MESSAGEPACK_TYPE.containing_type = _MESSAGEPACK
DESCRIPTOR.message_types_by_name['Frame'] = _FRAME
DESCRIPTOR.message_types_by_name['Check'] = _CHECK
DESCRIPTOR.message_types_by_name['Addition'] = _ADDITION
DESCRIPTOR.message_types_by_name['MessagePack'] = _MESSAGEPACK
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Frame = _reflection.GeneratedProtocolMessageType('Frame', (_message.Message,), {
  'DESCRIPTOR' : _FRAME,
  '__module__' : 'pack_pb2'
  # @@protoc_insertion_point(class_scope:pack.Frame)
  })
_sym_db.RegisterMessage(Frame)

Check = _reflection.GeneratedProtocolMessageType('Check', (_message.Message,), {
  'DESCRIPTOR' : _CHECK,
  '__module__' : 'pack_pb2'
  # @@protoc_insertion_point(class_scope:pack.Check)
  })
_sym_db.RegisterMessage(Check)

Addition = _reflection.GeneratedProtocolMessageType('Addition', (_message.Message,), {
  'DESCRIPTOR' : _ADDITION,
  '__module__' : 'pack_pb2'
  # @@protoc_insertion_point(class_scope:pack.Addition)
  })
_sym_db.RegisterMessage(Addition)

MessagePack = _reflection.GeneratedProtocolMessageType('MessagePack', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGEPACK,
  '__module__' : 'pack_pb2'
  # @@protoc_insertion_point(class_scope:pack.MessagePack)
  })
_sym_db.RegisterMessage(MessagePack)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
