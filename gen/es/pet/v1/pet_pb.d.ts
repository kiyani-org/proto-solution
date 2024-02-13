// @generated by protoc-gen-es v1.7.1
// @generated from file pet/v1/pet.proto (package pet.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * PetType represents the different types of pets in the pet store.
 *
 * @generated from enum pet.v1.PetType
 */
export declare enum PetType {
  /**
   * @generated from enum value: PET_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PET_TYPE_CAT = 1;
   */
  CAT = 1,

  /**
   * @generated from enum value: PET_TYPE_DOG = 2;
   */
  DOG = 2,

  /**
   * @generated from enum value: PET_TYPE_SNAKE = 3;
   */
  SNAKE = 3,

  /**
   * @generated from enum value: PET_TYPE_HAMSTER = 4;
   */
  HAMSTER = 4,
}

/**
 * Pet represents a pet in the pet store.
 *
 * @generated from message pet.v1.Pet
 */
export declare class Pet extends Message<Pet> {
  /**
   * @generated from field: pet.v1.PetType pet_type = 1;
   */
  petType: PetType;

  /**
   * @generated from field: string pet_id = 2;
   */
  petId: string;

  /**
   * @generated from field: string name = 3;
   */
  name: string;

  constructor(data?: PartialMessage<Pet>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "pet.v1.Pet";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pet;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pet;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pet;

  static equals(a: Pet | PlainMessage<Pet> | undefined, b: Pet | PlainMessage<Pet> | undefined): boolean;
}

/**
 * @generated from message pet.v1.GetPetRequest
 */
export declare class GetPetRequest extends Message<GetPetRequest> {
  /**
   * @generated from field: string pet_id = 1;
   */
  petId: string;

  constructor(data?: PartialMessage<GetPetRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "pet.v1.GetPetRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPetRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPetRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPetRequest;

  static equals(a: GetPetRequest | PlainMessage<GetPetRequest> | undefined, b: GetPetRequest | PlainMessage<GetPetRequest> | undefined): boolean;
}

/**
 * @generated from message pet.v1.GetPetResponse
 */
export declare class GetPetResponse extends Message<GetPetResponse> {
  /**
   * @generated from field: pet.v1.Pet pet = 1;
   */
  pet?: Pet;

  constructor(data?: PartialMessage<GetPetResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "pet.v1.GetPetResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPetResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPetResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPetResponse;

  static equals(a: GetPetResponse | PlainMessage<GetPetResponse> | undefined, b: GetPetResponse | PlainMessage<GetPetResponse> | undefined): boolean;
}

/**
 * @generated from message pet.v1.PutPetRequest
 */
export declare class PutPetRequest extends Message<PutPetRequest> {
  /**
   * @generated from field: pet.v1.PetType pet_type = 1;
   */
  petType: PetType;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  constructor(data?: PartialMessage<PutPetRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "pet.v1.PutPetRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PutPetRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PutPetRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PutPetRequest;

  static equals(a: PutPetRequest | PlainMessage<PutPetRequest> | undefined, b: PutPetRequest | PlainMessage<PutPetRequest> | undefined): boolean;
}

/**
 * @generated from message pet.v1.PutPetResponse
 */
export declare class PutPetResponse extends Message<PutPetResponse> {
  /**
   * @generated from field: pet.v1.Pet pet = 1;
   */
  pet?: Pet;

  constructor(data?: PartialMessage<PutPetResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "pet.v1.PutPetResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PutPetResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PutPetResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PutPetResponse;

  static equals(a: PutPetResponse | PlainMessage<PutPetResponse> | undefined, b: PutPetResponse | PlainMessage<PutPetResponse> | undefined): boolean;
}

/**
 * @generated from message pet.v1.DeletePetRequest
 */
export declare class DeletePetRequest extends Message<DeletePetRequest> {
  /**
   * @generated from field: string pet_id = 1;
   */
  petId: string;

  constructor(data?: PartialMessage<DeletePetRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "pet.v1.DeletePetRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeletePetRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeletePetRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeletePetRequest;

  static equals(a: DeletePetRequest | PlainMessage<DeletePetRequest> | undefined, b: DeletePetRequest | PlainMessage<DeletePetRequest> | undefined): boolean;
}

/**
 * @generated from message pet.v1.DeletePetResponse
 */
export declare class DeletePetResponse extends Message<DeletePetResponse> {
  constructor(data?: PartialMessage<DeletePetResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "pet.v1.DeletePetResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeletePetResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeletePetResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeletePetResponse;

  static equals(a: DeletePetResponse | PlainMessage<DeletePetResponse> | undefined, b: DeletePetResponse | PlainMessage<DeletePetResponse> | undefined): boolean;
}
