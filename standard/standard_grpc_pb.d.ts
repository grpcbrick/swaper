// GENERATED CODE -- DO NOT EDIT!

// package: standard
// file: standard.proto

import * as standard_pb from "./standard_pb";
import * as grpc from "grpc";

interface ISwitchService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  createUnitData: grpc.MethodDefinition<standard_pb.CreateUnitDataRequest, standard_pb.CreateUnitDataResponse>;
  getUnitDataByKey: grpc.MethodDefinition<standard_pb.GetUnitDataByKeyRequest, standard_pb.GetUnitDataByKeyResponse>;
  destroyUnitDataByKey: grpc.MethodDefinition<standard_pb.DestroyUnitDataByKeyRequest, standard_pb.DestroyUnitDataByKeyResponse>;
  updateUnitDataBodyByKey: grpc.MethodDefinition<standard_pb.UpdateUnitDataBodyByKeyRequest, standard_pb.UpdateUnitDataBodyByKeyResponse>;
  updateUnitDataExpiryTimeByKey: grpc.MethodDefinition<standard_pb.UpdateUnitDataExpiryTimeByKeyRequest, standard_pb.UpdateUnitDataExpiryTimeByKeyResponse>;
  updateUnitDataDestroyTimeByKey: grpc.MethodDefinition<standard_pb.UpdateUnitDataDestroyTimeByKeyRequest, standard_pb.UpdateUnitDataDestroyTimeByKeyResponse>;
  updateUnitDataEffectiveTimeByKey: grpc.MethodDefinition<standard_pb.UpdateUnitDataEffectiveTimeByKeyRequest, standard_pb.UpdateUnitDataEffectiveTimeByKeyResponse>;
}

export const SwitchService: ISwitchService;

export class SwitchClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  createUnitData(argument: standard_pb.CreateUnitDataRequest, callback: grpc.requestCallback<standard_pb.CreateUnitDataResponse>): grpc.ClientUnaryCall;
  createUnitData(argument: standard_pb.CreateUnitDataRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.CreateUnitDataResponse>): grpc.ClientUnaryCall;
  createUnitData(argument: standard_pb.CreateUnitDataRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.CreateUnitDataResponse>): grpc.ClientUnaryCall;
  getUnitDataByKey(argument: standard_pb.GetUnitDataByKeyRequest, callback: grpc.requestCallback<standard_pb.GetUnitDataByKeyResponse>): grpc.ClientUnaryCall;
  getUnitDataByKey(argument: standard_pb.GetUnitDataByKeyRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.GetUnitDataByKeyResponse>): grpc.ClientUnaryCall;
  getUnitDataByKey(argument: standard_pb.GetUnitDataByKeyRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.GetUnitDataByKeyResponse>): grpc.ClientUnaryCall;
  destroyUnitDataByKey(argument: standard_pb.DestroyUnitDataByKeyRequest, callback: grpc.requestCallback<standard_pb.DestroyUnitDataByKeyResponse>): grpc.ClientUnaryCall;
  destroyUnitDataByKey(argument: standard_pb.DestroyUnitDataByKeyRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DestroyUnitDataByKeyResponse>): grpc.ClientUnaryCall;
  destroyUnitDataByKey(argument: standard_pb.DestroyUnitDataByKeyRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.DestroyUnitDataByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataBodyByKey(argument: standard_pb.UpdateUnitDataBodyByKeyRequest, callback: grpc.requestCallback<standard_pb.UpdateUnitDataBodyByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataBodyByKey(argument: standard_pb.UpdateUnitDataBodyByKeyRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataBodyByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataBodyByKey(argument: standard_pb.UpdateUnitDataBodyByKeyRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataBodyByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataExpiryTimeByKey(argument: standard_pb.UpdateUnitDataExpiryTimeByKeyRequest, callback: grpc.requestCallback<standard_pb.UpdateUnitDataExpiryTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataExpiryTimeByKey(argument: standard_pb.UpdateUnitDataExpiryTimeByKeyRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataExpiryTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataExpiryTimeByKey(argument: standard_pb.UpdateUnitDataExpiryTimeByKeyRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataExpiryTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataDestroyTimeByKey(argument: standard_pb.UpdateUnitDataDestroyTimeByKeyRequest, callback: grpc.requestCallback<standard_pb.UpdateUnitDataDestroyTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataDestroyTimeByKey(argument: standard_pb.UpdateUnitDataDestroyTimeByKeyRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataDestroyTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataDestroyTimeByKey(argument: standard_pb.UpdateUnitDataDestroyTimeByKeyRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataDestroyTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataEffectiveTimeByKey(argument: standard_pb.UpdateUnitDataEffectiveTimeByKeyRequest, callback: grpc.requestCallback<standard_pb.UpdateUnitDataEffectiveTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataEffectiveTimeByKey(argument: standard_pb.UpdateUnitDataEffectiveTimeByKeyRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataEffectiveTimeByKeyResponse>): grpc.ClientUnaryCall;
  updateUnitDataEffectiveTimeByKey(argument: standard_pb.UpdateUnitDataEffectiveTimeByKeyRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<standard_pb.UpdateUnitDataEffectiveTimeByKeyResponse>): grpc.ClientUnaryCall;
}
