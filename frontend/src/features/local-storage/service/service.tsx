export const LocalStorage = {
    get(key:string):string   {
        let value = localStorage.getItem(key)
        if (value === null){
            return ""
        }
        return value

    },
    set(key:string, value:string) {
        localStorage.setItem(key, value)
    },
    remove(key:string) {
        localStorage.removeItem(key)
    },
    clear() {
        localStorage.clear()
    },
  }