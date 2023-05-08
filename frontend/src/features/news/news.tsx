import { useEffect, useState } from "react";
import { getRequest } from "../../api/api";

type AlbumsResponse = {
    access_token:string 
    refresh_token:string
    logged_in:string
  }

export const News: React.FC = () => {

    const [albums,setAlbums] = useState()

    const getLatestAlbums = async (url:string) =>{
        return getRequest("auth/login").then(r => r.data).then(a => setAlbums(a))
        .catch(error => {   
          console.log(error)
        });
      }
    
    useEffect(() => {
        getLatestAlbums("")
    }, []);
    
    return (
        <>
        </>
    )
}