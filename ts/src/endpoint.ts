import type { Data } from './data'
import type { ObjectValues } from './generic'

//Map of available Options
const HTTPMethods = {
  POST: "POST",
  GET: "GET",
  PUT: "PUT",
  DELETE: "DELETE",
  OPTIONS: "OPTIONS",
  PATCH: "PATCH",
} as const

type HTTPMethod = ObjectValues<typeof HTTPMethods>

type endpoint = {
  method: HTTPMethod;
  response: Data;
  body?: Data;
  path: string;
  description?: string;
  error?: Data;
  params?: Data;
}
