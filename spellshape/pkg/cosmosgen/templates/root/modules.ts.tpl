import { SpellshapeClient } from "./client";
import { GeneratedType } from "@cosmjs/proto-signing";

export type ModuleInterface = { [key: string]: any }
export type Module = (instance: SpellshapeClient) => { module: ModuleInterface, registry: [string, GeneratedType][] }
