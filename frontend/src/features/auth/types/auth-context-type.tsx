import { Dispatch, SetStateAction } from "react"

export type AuthContextType = {
  auth: string
  setAuth: Dispatch<SetStateAction<string>>
  // add here other variables and functions 
  // you want to export with your context
}