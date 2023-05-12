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
        return getRequest("music/albums?limit=10&sort_by=created_at").then(r => console.log(r))
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