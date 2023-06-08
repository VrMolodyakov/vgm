export type Info = {
  barcode: string
  catalog_number: string
  classification: string
  currency_code: string
  full_image_src: string
  small_image_src: string
  media_format: string
  price: number
  publisher: string
}

export type Track = {
  title: string
  duration: string
}

export type Credit = {
  profession: string
  person_id: number
}

export type Album = {
  title: string
  released_at: Number
} 

export type FullAlbum = {
  album: Album
  info: Info
  credits: Credit[]
  tracklist: Track[]
}