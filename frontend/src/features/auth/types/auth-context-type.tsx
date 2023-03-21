export type AuthContextType = {
  auth: string | null
  setAuth: React.Dispatch<React.SetStateAction<string | null>>
}