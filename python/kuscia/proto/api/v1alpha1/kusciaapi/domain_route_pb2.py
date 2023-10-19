# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/kusciaapi/domain_route.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n6kuscia/proto/api/v1alpha1/kusciaapi/domain_route.proto\x12#kuscia.proto.api.v1alpha1.kusciaapi\x1a&kuscia/proto/api/v1alpha1/common.proto\"\xea\x02\n\x18\x43reateDomainRouteRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x1b\n\x13\x61uthentication_type\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65stination\x18\x03 \x01(\t\x12\x44\n\x08\x65ndpoint\x18\x04 \x01(\x0b\x32\x32.kuscia.proto.api.v1alpha1.kusciaapi.RouteEndpoint\x12\x0e\n\x06source\x18\x05 \x01(\t\x12\x46\n\x0ctoken_config\x18\x06 \x01(\x0b\x32\x30.kuscia.proto.api.v1alpha1.kusciaapi.TokenConfig\x12\x44\n\x0bmtls_config\x18\x07 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.MTLSConfig\"_\n\rRouteEndpoint\x12\x0c\n\x04host\x18\x01 \x01(\t\x12@\n\x05ports\x18\x02 \x03(\x0b\x32\x31.kuscia.proto.api.v1alpha1.kusciaapi.EndpointPort\"<\n\x0c\x45ndpointPort\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04port\x18\x02 \x01(\x05\x12\x10\n\x08protocol\x18\x03 \x01(\t\"\x81\x01\n\x0bTokenConfig\x12\x1e\n\x16\x64\x65stination_public_key\x18\x01 \x01(\t\x12\x1d\n\x15rolling_update_period\x18\x02 \x01(\x03\x12\x19\n\x11source_public_key\x18\x03 \x01(\t\x12\x18\n\x10token_gen_method\x18\x04 \x01(\t\"[\n\nMTLSConfig\x12\x0e\n\x06tls_ca\x18\x01 \x01(\t\x12!\n\x19source_client_private_key\x18\x02 \x01(\t\x12\x1a\n\x12source_client_cert\x18\x03 \x01(\t\"\xa0\x01\n\x19\x43reateDomainRouteResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12P\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x42.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainRouteResponseData\"-\n\x1d\x43reateDomainRouteResponseData\x12\x0c\n\x04name\x18\x01 \x01(\t\"y\n\x18\x44\x65leteDomainRouteRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x13\n\x0b\x64\x65stination\x18\x03 \x01(\t\x12\x0e\n\x06source\x18\x02 \x01(\t\"N\n\x19\x44\x65leteDomainRouteResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"x\n\x17QueryDomainRouteRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x13\n\x0b\x64\x65stination\x18\x02 \x01(\t\x12\x0e\n\x06source\x18\x03 \x01(\t\"\x9e\x01\n\x18QueryDomainRouteResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12O\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x41.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainRouteResponseData\"\x84\x03\n\x1cQueryDomainRouteResponseData\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x1b\n\x13\x61uthentication_type\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65stination\x18\x03 \x01(\t\x12\x44\n\x08\x65ndpoint\x18\x04 \x01(\x0b\x32\x32.kuscia.proto.api.v1alpha1.kusciaapi.RouteEndpoint\x12\x0e\n\x06source\x18\x05 \x01(\t\x12\x46\n\x0ctoken_config\x18\x06 \x01(\x0b\x32\x30.kuscia.proto.api.v1alpha1.kusciaapi.TokenConfig\x12\x44\n\x0bmtls_config\x18\x07 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.MTLSConfig\x12@\n\x06status\x18\x08 \x01(\x0b\x32\x30.kuscia.proto.api.v1alpha1.kusciaapi.RouteStatus\"-\n\x0bRouteStatus\x12\x0e\n\x06status\x18\x01 \x01(\t\x12\x0e\n\x06reason\x18\x02 \x01(\t\"\xa7\x01\n\"BatchQueryDomainRouteStatusRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12G\n\nroute_keys\x18\x02 \x03(\x0b\x32\x33.kuscia.proto.api.v1alpha1.kusciaapi.DomainRouteKey\"5\n\x0e\x44omainRouteKey\x12\x0e\n\x06source\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65stination\x18\x02 \x01(\t\"\xb4\x01\n#BatchQueryDomainRouteStatusResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12Z\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32L.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainRouteStatusResponseData\"q\n\'BatchQueryDomainRouteStatusResponseData\x12\x46\n\x06routes\x18\x01 \x03(\x0b\x32\x36.kuscia.proto.api.v1alpha1.kusciaapi.DomainRouteStatus\"\x88\x01\n\x11\x44omainRouteStatus\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65stination\x18\x02 \x01(\t\x12\x0e\n\x06source\x18\x03 \x01(\t\x12@\n\x06status\x18\x04 \x01(\x0b\x32\x30.kuscia.proto.api.v1alpha1.kusciaapi.RouteStatus*)\n\x12\x41uthenticationType\x12\t\n\x05Token\x10\x00\x12\x08\n\x04MTLS\x10\x01\x32\x83\x05\n\x12\x44omainRouteService\x12\x92\x01\n\x11\x43reateDomainRoute\x12=.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainRouteRequest\x1a>.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainRouteResponse\x12\x92\x01\n\x11\x44\x65leteDomainRoute\x12=.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainRouteRequest\x1a>.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainRouteResponse\x12\x8f\x01\n\x10QueryDomainRoute\x12<.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainRouteRequest\x1a=.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainRouteResponse\x12\xb0\x01\n\x1b\x42\x61tchQueryDomainRouteStatus\x12G.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainRouteStatusRequest\x1aH.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainRouteStatusResponseB^\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapib\x06proto3')

_AUTHENTICATIONTYPE = DESCRIPTOR.enum_types_by_name['AuthenticationType']
AuthenticationType = enum_type_wrapper.EnumTypeWrapper(_AUTHENTICATIONTYPE)
Token = 0
MTLS = 1


_CREATEDOMAINROUTEREQUEST = DESCRIPTOR.message_types_by_name['CreateDomainRouteRequest']
_ROUTEENDPOINT = DESCRIPTOR.message_types_by_name['RouteEndpoint']
_ENDPOINTPORT = DESCRIPTOR.message_types_by_name['EndpointPort']
_TOKENCONFIG = DESCRIPTOR.message_types_by_name['TokenConfig']
_MTLSCONFIG = DESCRIPTOR.message_types_by_name['MTLSConfig']
_CREATEDOMAINROUTERESPONSE = DESCRIPTOR.message_types_by_name['CreateDomainRouteResponse']
_CREATEDOMAINROUTERESPONSEDATA = DESCRIPTOR.message_types_by_name['CreateDomainRouteResponseData']
_DELETEDOMAINROUTEREQUEST = DESCRIPTOR.message_types_by_name['DeleteDomainRouteRequest']
_DELETEDOMAINROUTERESPONSE = DESCRIPTOR.message_types_by_name['DeleteDomainRouteResponse']
_QUERYDOMAINROUTEREQUEST = DESCRIPTOR.message_types_by_name['QueryDomainRouteRequest']
_QUERYDOMAINROUTERESPONSE = DESCRIPTOR.message_types_by_name['QueryDomainRouteResponse']
_QUERYDOMAINROUTERESPONSEDATA = DESCRIPTOR.message_types_by_name['QueryDomainRouteResponseData']
_ROUTESTATUS = DESCRIPTOR.message_types_by_name['RouteStatus']
_BATCHQUERYDOMAINROUTESTATUSREQUEST = DESCRIPTOR.message_types_by_name['BatchQueryDomainRouteStatusRequest']
_DOMAINROUTEKEY = DESCRIPTOR.message_types_by_name['DomainRouteKey']
_BATCHQUERYDOMAINROUTESTATUSRESPONSE = DESCRIPTOR.message_types_by_name['BatchQueryDomainRouteStatusResponse']
_BATCHQUERYDOMAINROUTESTATUSRESPONSEDATA = DESCRIPTOR.message_types_by_name['BatchQueryDomainRouteStatusResponseData']
_DOMAINROUTESTATUS = DESCRIPTOR.message_types_by_name['DomainRouteStatus']
CreateDomainRouteRequest = _reflection.GeneratedProtocolMessageType('CreateDomainRouteRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDOMAINROUTEREQUEST,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainRouteRequest)
  })
_sym_db.RegisterMessage(CreateDomainRouteRequest)

RouteEndpoint = _reflection.GeneratedProtocolMessageType('RouteEndpoint', (_message.Message,), {
  'DESCRIPTOR' : _ROUTEENDPOINT,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.RouteEndpoint)
  })
_sym_db.RegisterMessage(RouteEndpoint)

EndpointPort = _reflection.GeneratedProtocolMessageType('EndpointPort', (_message.Message,), {
  'DESCRIPTOR' : _ENDPOINTPORT,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.EndpointPort)
  })
_sym_db.RegisterMessage(EndpointPort)

TokenConfig = _reflection.GeneratedProtocolMessageType('TokenConfig', (_message.Message,), {
  'DESCRIPTOR' : _TOKENCONFIG,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.TokenConfig)
  })
_sym_db.RegisterMessage(TokenConfig)

MTLSConfig = _reflection.GeneratedProtocolMessageType('MTLSConfig', (_message.Message,), {
  'DESCRIPTOR' : _MTLSCONFIG,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.MTLSConfig)
  })
_sym_db.RegisterMessage(MTLSConfig)

CreateDomainRouteResponse = _reflection.GeneratedProtocolMessageType('CreateDomainRouteResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDOMAINROUTERESPONSE,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainRouteResponse)
  })
_sym_db.RegisterMessage(CreateDomainRouteResponse)

CreateDomainRouteResponseData = _reflection.GeneratedProtocolMessageType('CreateDomainRouteResponseData', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDOMAINROUTERESPONSEDATA,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainRouteResponseData)
  })
_sym_db.RegisterMessage(CreateDomainRouteResponseData)

DeleteDomainRouteRequest = _reflection.GeneratedProtocolMessageType('DeleteDomainRouteRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEDOMAINROUTEREQUEST,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainRouteRequest)
  })
_sym_db.RegisterMessage(DeleteDomainRouteRequest)

DeleteDomainRouteResponse = _reflection.GeneratedProtocolMessageType('DeleteDomainRouteResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETEDOMAINROUTERESPONSE,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainRouteResponse)
  })
_sym_db.RegisterMessage(DeleteDomainRouteResponse)

QueryDomainRouteRequest = _reflection.GeneratedProtocolMessageType('QueryDomainRouteRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUERYDOMAINROUTEREQUEST,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainRouteRequest)
  })
_sym_db.RegisterMessage(QueryDomainRouteRequest)

QueryDomainRouteResponse = _reflection.GeneratedProtocolMessageType('QueryDomainRouteResponse', (_message.Message,), {
  'DESCRIPTOR' : _QUERYDOMAINROUTERESPONSE,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainRouteResponse)
  })
_sym_db.RegisterMessage(QueryDomainRouteResponse)

QueryDomainRouteResponseData = _reflection.GeneratedProtocolMessageType('QueryDomainRouteResponseData', (_message.Message,), {
  'DESCRIPTOR' : _QUERYDOMAINROUTERESPONSEDATA,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainRouteResponseData)
  })
_sym_db.RegisterMessage(QueryDomainRouteResponseData)

RouteStatus = _reflection.GeneratedProtocolMessageType('RouteStatus', (_message.Message,), {
  'DESCRIPTOR' : _ROUTESTATUS,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.RouteStatus)
  })
_sym_db.RegisterMessage(RouteStatus)

BatchQueryDomainRouteStatusRequest = _reflection.GeneratedProtocolMessageType('BatchQueryDomainRouteStatusRequest', (_message.Message,), {
  'DESCRIPTOR' : _BATCHQUERYDOMAINROUTESTATUSREQUEST,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainRouteStatusRequest)
  })
_sym_db.RegisterMessage(BatchQueryDomainRouteStatusRequest)

DomainRouteKey = _reflection.GeneratedProtocolMessageType('DomainRouteKey', (_message.Message,), {
  'DESCRIPTOR' : _DOMAINROUTEKEY,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.DomainRouteKey)
  })
_sym_db.RegisterMessage(DomainRouteKey)

BatchQueryDomainRouteStatusResponse = _reflection.GeneratedProtocolMessageType('BatchQueryDomainRouteStatusResponse', (_message.Message,), {
  'DESCRIPTOR' : _BATCHQUERYDOMAINROUTESTATUSRESPONSE,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainRouteStatusResponse)
  })
_sym_db.RegisterMessage(BatchQueryDomainRouteStatusResponse)

BatchQueryDomainRouteStatusResponseData = _reflection.GeneratedProtocolMessageType('BatchQueryDomainRouteStatusResponseData', (_message.Message,), {
  'DESCRIPTOR' : _BATCHQUERYDOMAINROUTESTATUSRESPONSEDATA,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainRouteStatusResponseData)
  })
_sym_db.RegisterMessage(BatchQueryDomainRouteStatusResponseData)

DomainRouteStatus = _reflection.GeneratedProtocolMessageType('DomainRouteStatus', (_message.Message,), {
  'DESCRIPTOR' : _DOMAINROUTESTATUS,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.domain_route_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.DomainRouteStatus)
  })
_sym_db.RegisterMessage(DomainRouteStatus)

_DOMAINROUTESERVICE = DESCRIPTOR.services_by_name['DomainRouteService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi'
  _AUTHENTICATIONTYPE._serialized_start=2680
  _AUTHENTICATIONTYPE._serialized_end=2721
  _CREATEDOMAINROUTEREQUEST._serialized_start=136
  _CREATEDOMAINROUTEREQUEST._serialized_end=498
  _ROUTEENDPOINT._serialized_start=500
  _ROUTEENDPOINT._serialized_end=595
  _ENDPOINTPORT._serialized_start=597
  _ENDPOINTPORT._serialized_end=657
  _TOKENCONFIG._serialized_start=660
  _TOKENCONFIG._serialized_end=789
  _MTLSCONFIG._serialized_start=791
  _MTLSCONFIG._serialized_end=882
  _CREATEDOMAINROUTERESPONSE._serialized_start=885
  _CREATEDOMAINROUTERESPONSE._serialized_end=1045
  _CREATEDOMAINROUTERESPONSEDATA._serialized_start=1047
  _CREATEDOMAINROUTERESPONSEDATA._serialized_end=1092
  _DELETEDOMAINROUTEREQUEST._serialized_start=1094
  _DELETEDOMAINROUTEREQUEST._serialized_end=1215
  _DELETEDOMAINROUTERESPONSE._serialized_start=1217
  _DELETEDOMAINROUTERESPONSE._serialized_end=1295
  _QUERYDOMAINROUTEREQUEST._serialized_start=1297
  _QUERYDOMAINROUTEREQUEST._serialized_end=1417
  _QUERYDOMAINROUTERESPONSE._serialized_start=1420
  _QUERYDOMAINROUTERESPONSE._serialized_end=1578
  _QUERYDOMAINROUTERESPONSEDATA._serialized_start=1581
  _QUERYDOMAINROUTERESPONSEDATA._serialized_end=1969
  _ROUTESTATUS._serialized_start=1971
  _ROUTESTATUS._serialized_end=2016
  _BATCHQUERYDOMAINROUTESTATUSREQUEST._serialized_start=2019
  _BATCHQUERYDOMAINROUTESTATUSREQUEST._serialized_end=2186
  _DOMAINROUTEKEY._serialized_start=2188
  _DOMAINROUTEKEY._serialized_end=2241
  _BATCHQUERYDOMAINROUTESTATUSRESPONSE._serialized_start=2244
  _BATCHQUERYDOMAINROUTESTATUSRESPONSE._serialized_end=2424
  _BATCHQUERYDOMAINROUTESTATUSRESPONSEDATA._serialized_start=2426
  _BATCHQUERYDOMAINROUTESTATUSRESPONSEDATA._serialized_end=2539
  _DOMAINROUTESTATUS._serialized_start=2542
  _DOMAINROUTESTATUS._serialized_end=2678
  _DOMAINROUTESERVICE._serialized_start=2724
  _DOMAINROUTESERVICE._serialized_end=3367
# @@protoc_insertion_point(module_scope)