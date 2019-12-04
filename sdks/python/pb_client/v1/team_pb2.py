#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/team.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='v1/team.proto',
  package='v1',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\rv1/team.proto\x12\x02v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\x94\x01\n\x04Team\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08projects\x18\x03 \x03(\t\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"8\n\x0fTeamBodyRequest\x12\r\n\x05owner\x18\x01 \x01(\t\x12\x16\n\x04team\x18\x02 \x01(\x0b\x32\x08.v1.Team\"]\n\x11ListTeamsResponse\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\x12\x19\n\x07results\x18\x02 \x03(\x0b\x32\x08.v1.Team\x12\x10\n\x08previous\x18\x03 \x01(\t\x12\x0c\n\x04next\x18\x04 \x01(\t\"\x9a\x01\n\nTeamMember\x12\x0c\n\x04user\x18\x01 \x01(\t\x12\x0c\n\x04role\x18\x02 \x01(\t\x12\x10\n\x08org_role\x18\x03 \x01(\t\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"T\n\x15TeamMemberBodyRequest\x12\r\n\x05owner\x18\x01 \x01(\t\x12\x0c\n\x04team\x18\x02 \x01(\t\x12\x1e\n\x06member\x18\x03 \x01(\x0b\x32\x0e.v1.TeamMember\"i\n\x17ListTeamMembersResponse\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\x12\x1f\n\x07results\x18\x02 \x03(\x0b\x32\x0e.v1.TeamMember\x12\x10\n\x08previous\x18\x03 \x01(\t\x12\x0c\n\x04next\x18\x04 \x01(\tb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_TEAM = _descriptor.Descriptor(
  name='Team',
  full_name='v1.Team',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='v1.Team.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='v1.Team.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='projects', full_name='v1.Team.projects', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='v1.Team.created_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='v1.Team.updated_at', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=55,
  serialized_end=203,
)


_TEAMBODYREQUEST = _descriptor.Descriptor(
  name='TeamBodyRequest',
  full_name='v1.TeamBodyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner', full_name='v1.TeamBodyRequest.owner', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='team', full_name='v1.TeamBodyRequest.team', index=1,
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
  serialized_start=205,
  serialized_end=261,
)


_LISTTEAMSRESPONSE = _descriptor.Descriptor(
  name='ListTeamsResponse',
  full_name='v1.ListTeamsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='count', full_name='v1.ListTeamsResponse.count', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='results', full_name='v1.ListTeamsResponse.results', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='previous', full_name='v1.ListTeamsResponse.previous', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next', full_name='v1.ListTeamsResponse.next', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=263,
  serialized_end=356,
)


_TEAMMEMBER = _descriptor.Descriptor(
  name='TeamMember',
  full_name='v1.TeamMember',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='v1.TeamMember.user', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='role', full_name='v1.TeamMember.role', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='org_role', full_name='v1.TeamMember.org_role', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='v1.TeamMember.created_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='v1.TeamMember.updated_at', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=359,
  serialized_end=513,
)


_TEAMMEMBERBODYREQUEST = _descriptor.Descriptor(
  name='TeamMemberBodyRequest',
  full_name='v1.TeamMemberBodyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner', full_name='v1.TeamMemberBodyRequest.owner', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='team', full_name='v1.TeamMemberBodyRequest.team', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='member', full_name='v1.TeamMemberBodyRequest.member', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=515,
  serialized_end=599,
)


_LISTTEAMMEMBERSRESPONSE = _descriptor.Descriptor(
  name='ListTeamMembersResponse',
  full_name='v1.ListTeamMembersResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='count', full_name='v1.ListTeamMembersResponse.count', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='results', full_name='v1.ListTeamMembersResponse.results', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='previous', full_name='v1.ListTeamMembersResponse.previous', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next', full_name='v1.ListTeamMembersResponse.next', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=601,
  serialized_end=706,
)

_TEAM.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TEAM.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TEAMBODYREQUEST.fields_by_name['team'].message_type = _TEAM
_LISTTEAMSRESPONSE.fields_by_name['results'].message_type = _TEAM
_TEAMMEMBER.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TEAMMEMBER.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TEAMMEMBERBODYREQUEST.fields_by_name['member'].message_type = _TEAMMEMBER
_LISTTEAMMEMBERSRESPONSE.fields_by_name['results'].message_type = _TEAMMEMBER
DESCRIPTOR.message_types_by_name['Team'] = _TEAM
DESCRIPTOR.message_types_by_name['TeamBodyRequest'] = _TEAMBODYREQUEST
DESCRIPTOR.message_types_by_name['ListTeamsResponse'] = _LISTTEAMSRESPONSE
DESCRIPTOR.message_types_by_name['TeamMember'] = _TEAMMEMBER
DESCRIPTOR.message_types_by_name['TeamMemberBodyRequest'] = _TEAMMEMBERBODYREQUEST
DESCRIPTOR.message_types_by_name['ListTeamMembersResponse'] = _LISTTEAMMEMBERSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Team = _reflection.GeneratedProtocolMessageType('Team', (_message.Message,), {
  'DESCRIPTOR' : _TEAM,
  '__module__' : 'v1.team_pb2'
  # @@protoc_insertion_point(class_scope:v1.Team)
  })
_sym_db.RegisterMessage(Team)

TeamBodyRequest = _reflection.GeneratedProtocolMessageType('TeamBodyRequest', (_message.Message,), {
  'DESCRIPTOR' : _TEAMBODYREQUEST,
  '__module__' : 'v1.team_pb2'
  # @@protoc_insertion_point(class_scope:v1.TeamBodyRequest)
  })
_sym_db.RegisterMessage(TeamBodyRequest)

ListTeamsResponse = _reflection.GeneratedProtocolMessageType('ListTeamsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTTEAMSRESPONSE,
  '__module__' : 'v1.team_pb2'
  # @@protoc_insertion_point(class_scope:v1.ListTeamsResponse)
  })
_sym_db.RegisterMessage(ListTeamsResponse)

TeamMember = _reflection.GeneratedProtocolMessageType('TeamMember', (_message.Message,), {
  'DESCRIPTOR' : _TEAMMEMBER,
  '__module__' : 'v1.team_pb2'
  # @@protoc_insertion_point(class_scope:v1.TeamMember)
  })
_sym_db.RegisterMessage(TeamMember)

TeamMemberBodyRequest = _reflection.GeneratedProtocolMessageType('TeamMemberBodyRequest', (_message.Message,), {
  'DESCRIPTOR' : _TEAMMEMBERBODYREQUEST,
  '__module__' : 'v1.team_pb2'
  # @@protoc_insertion_point(class_scope:v1.TeamMemberBodyRequest)
  })
_sym_db.RegisterMessage(TeamMemberBodyRequest)

ListTeamMembersResponse = _reflection.GeneratedProtocolMessageType('ListTeamMembersResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTTEAMMEMBERSRESPONSE,
  '__module__' : 'v1.team_pb2'
  # @@protoc_insertion_point(class_scope:v1.ListTeamMembersResponse)
  })
_sym_db.RegisterMessage(ListTeamMembersResponse)


# @@protoc_insertion_point(module_scope)
