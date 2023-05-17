export type AlbumView = {
    album_id:string
    created_at:number
    released_at:number
    title:string
    publisher:string
    small_image_src:string
}

export type Albums = {
    albums:Array<AlbumView>
}