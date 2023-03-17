export type AuthContextType = {
    auth:string | undefined,
    saveAuth:(auth:string) => void
}