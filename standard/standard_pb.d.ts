// package: standard
// file: standard.proto

import * as jspb from "google-protobuf";

export class UnitData extends jspb.Message {
  getBody(): string;
  setBody(value: string): void;

  getExpirytime(): string;
  setExpirytime(value: string): void;

  getCreatedtime(): string;
  setCreatedtime(value: string): void;

  getDestroytime(): string;
  setDestroytime(value: string): void;

  getEffectivetime(): string;
  setEffectivetime(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnitData.AsObject;
  static toObject(includeInstance: boolean, msg: UnitData): UnitData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnitData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnitData;
  static deserializeBinaryFromReader(message: UnitData, reader: jspb.BinaryReader): UnitData;
}

export namespace UnitData {
  export type AsObject = {
    body: string,
    expirytime: string,
    createdtime: string,
    destroytime: string,
    effectivetime: string,
  }
}

export class CreateUnitDataRequest extends jspb.Message {
  getBody(): string;
  setBody(value: string): void;

  getExpirytime(): string;
  setExpirytime(value: string): void;

  getDestroytime(): string;
  setDestroytime(value: string): void;

  getEffectivetime(): string;
  setEffectivetime(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUnitDataRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUnitDataRequest): CreateUnitDataRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUnitDataRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUnitDataRequest;
  static deserializeBinaryFromReader(message: CreateUnitDataRequest, reader: jspb.BinaryReader): CreateUnitDataRequest;
}

export namespace CreateUnitDataRequest {
  export type AsObject = {
    body: string,
    expirytime: string,
    destroytime: string,
    effectivetime: string,
  }
}

export class CreateUnitDataResponse extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUnitDataResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUnitDataResponse): CreateUnitDataResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUnitDataResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUnitDataResponse;
  static deserializeBinaryFromReader(message: CreateUnitDataResponse, reader: jspb.BinaryReader): CreateUnitDataResponse;
}

export namespace CreateUnitDataResponse {
  export type AsObject = {
    key: string,
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class GetUnitDataByKeyRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUnitDataByKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUnitDataByKeyRequest): GetUnitDataByKeyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUnitDataByKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUnitDataByKeyRequest;
  static deserializeBinaryFromReader(message: GetUnitDataByKeyRequest, reader: jspb.BinaryReader): GetUnitDataByKeyRequest;
}

export namespace GetUnitDataByKeyRequest {
  export type AsObject = {
    key: string,
  }
}

export class GetUnitDataByKeyResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasData(): boolean;
  clearData(): void;
  getData(): UnitData | undefined;
  setData(value?: UnitData): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUnitDataByKeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUnitDataByKeyResponse): GetUnitDataByKeyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUnitDataByKeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUnitDataByKeyResponse;
  static deserializeBinaryFromReader(message: GetUnitDataByKeyResponse, reader: jspb.BinaryReader): GetUnitDataByKeyResponse;
}

export namespace GetUnitDataByKeyResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
    data?: UnitData.AsObject,
  }
}

export class DestroyUnitDataByKeyRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyUnitDataByKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyUnitDataByKeyRequest): DestroyUnitDataByKeyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DestroyUnitDataByKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyUnitDataByKeyRequest;
  static deserializeBinaryFromReader(message: DestroyUnitDataByKeyRequest, reader: jspb.BinaryReader): DestroyUnitDataByKeyRequest;
}

export namespace DestroyUnitDataByKeyRequest {
  export type AsObject = {
    key: string,
  }
}

export class DestroyUnitDataByKeyResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyUnitDataByKeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyUnitDataByKeyResponse): DestroyUnitDataByKeyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DestroyUnitDataByKeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyUnitDataByKeyResponse;
  static deserializeBinaryFromReader(message: DestroyUnitDataByKeyResponse, reader: jspb.BinaryReader): DestroyUnitDataByKeyResponse;
}

export namespace DestroyUnitDataByKeyResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateUnitDataBodyByKeyRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataBodyByKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataBodyByKeyRequest): UpdateUnitDataBodyByKeyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataBodyByKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataBodyByKeyRequest;
  static deserializeBinaryFromReader(message: UpdateUnitDataBodyByKeyRequest, reader: jspb.BinaryReader): UpdateUnitDataBodyByKeyRequest;
}

export namespace UpdateUnitDataBodyByKeyRequest {
  export type AsObject = {
    key: string,
    body: string,
  }
}

export class UpdateUnitDataBodyByKeyResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataBodyByKeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataBodyByKeyResponse): UpdateUnitDataBodyByKeyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataBodyByKeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataBodyByKeyResponse;
  static deserializeBinaryFromReader(message: UpdateUnitDataBodyByKeyResponse, reader: jspb.BinaryReader): UpdateUnitDataBodyByKeyResponse;
}

export namespace UpdateUnitDataBodyByKeyResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateUnitDataExpiryTimeByKeyRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getExpirytime(): string;
  setExpirytime(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataExpiryTimeByKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataExpiryTimeByKeyRequest): UpdateUnitDataExpiryTimeByKeyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataExpiryTimeByKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataExpiryTimeByKeyRequest;
  static deserializeBinaryFromReader(message: UpdateUnitDataExpiryTimeByKeyRequest, reader: jspb.BinaryReader): UpdateUnitDataExpiryTimeByKeyRequest;
}

export namespace UpdateUnitDataExpiryTimeByKeyRequest {
  export type AsObject = {
    key: string,
    expirytime: string,
  }
}

export class UpdateUnitDataExpiryTimeByKeyResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataExpiryTimeByKeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataExpiryTimeByKeyResponse): UpdateUnitDataExpiryTimeByKeyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataExpiryTimeByKeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataExpiryTimeByKeyResponse;
  static deserializeBinaryFromReader(message: UpdateUnitDataExpiryTimeByKeyResponse, reader: jspb.BinaryReader): UpdateUnitDataExpiryTimeByKeyResponse;
}

export namespace UpdateUnitDataExpiryTimeByKeyResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateUnitDataDestroyTimeByKeyRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getDestroytime(): string;
  setDestroytime(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataDestroyTimeByKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataDestroyTimeByKeyRequest): UpdateUnitDataDestroyTimeByKeyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataDestroyTimeByKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataDestroyTimeByKeyRequest;
  static deserializeBinaryFromReader(message: UpdateUnitDataDestroyTimeByKeyRequest, reader: jspb.BinaryReader): UpdateUnitDataDestroyTimeByKeyRequest;
}

export namespace UpdateUnitDataDestroyTimeByKeyRequest {
  export type AsObject = {
    key: string,
    destroytime: string,
  }
}

export class UpdateUnitDataDestroyTimeByKeyResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataDestroyTimeByKeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataDestroyTimeByKeyResponse): UpdateUnitDataDestroyTimeByKeyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataDestroyTimeByKeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataDestroyTimeByKeyResponse;
  static deserializeBinaryFromReader(message: UpdateUnitDataDestroyTimeByKeyResponse, reader: jspb.BinaryReader): UpdateUnitDataDestroyTimeByKeyResponse;
}

export namespace UpdateUnitDataDestroyTimeByKeyResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export class UpdateUnitDataEffectiveTimeByKeyRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getEffectivetime(): string;
  setEffectivetime(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataEffectiveTimeByKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataEffectiveTimeByKeyRequest): UpdateUnitDataEffectiveTimeByKeyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataEffectiveTimeByKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataEffectiveTimeByKeyRequest;
  static deserializeBinaryFromReader(message: UpdateUnitDataEffectiveTimeByKeyRequest, reader: jspb.BinaryReader): UpdateUnitDataEffectiveTimeByKeyRequest;
}

export namespace UpdateUnitDataEffectiveTimeByKeyRequest {
  export type AsObject = {
    key: string,
    effectivetime: string,
  }
}

export class UpdateUnitDataEffectiveTimeByKeyResponse extends jspb.Message {
  getState(): StateMap[keyof StateMap];
  setState(value: StateMap[keyof StateMap]): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUnitDataEffectiveTimeByKeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUnitDataEffectiveTimeByKeyResponse): UpdateUnitDataEffectiveTimeByKeyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUnitDataEffectiveTimeByKeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUnitDataEffectiveTimeByKeyResponse;
  static deserializeBinaryFromReader(message: UpdateUnitDataEffectiveTimeByKeyResponse, reader: jspb.BinaryReader): UpdateUnitDataEffectiveTimeByKeyResponse;
}

export namespace UpdateUnitDataEffectiveTimeByKeyResponse {
  export type AsObject = {
    state: StateMap[keyof StateMap],
    message: string,
  }
}

export interface StateMap {
  UNKNOWN: 0;
  SUCCESS: 1;
  FAILURE: 2;
  EXPIRED: 3;
  NOT_ACTIVE: 4;
  SERVICE_ERROR: 5;
  PARAMS_INVALID: 6;
  ILLEGAL_REQUEST: 7;
  DB_OPERATION_FATLURE: 8;
}

export const State: StateMap;

