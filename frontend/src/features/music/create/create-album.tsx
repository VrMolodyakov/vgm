import { useForm } from "react-hook-form";
import { ChangeEvent, useState } from "react"
import "./create-album.css"
import DateInput from "../../../components/date-input/date-input";
import { Info, Track, Credit, Album, FullAlbum } from "./types";
import { useAlbum } from "./hooks/use-album";
import { MusicService } from "../service/music";
import { useMusicClient } from "../client-provider/context/context";

const CreateForm: React.FC = () => {
  const { register, handleSubmit } = useForm<Album & Info>();
  const [date, setDate] = useState(new Date())
  const [tracklist, setTracklist] = useState<Track[]>([])
  const [credits, setCredits] = useState<Credit[]>([])
  let client = useMusicClient()
  let musicService = new MusicService(client)
  const { data, error, mutate: create, isSuccess, isError } = useAlbum(musicService)

  async function onSubmit(data: Album & Info){
    const album:Album = {
      title:data.title,
      released_at:new Number(date),
    }

    const info:Info = {
      barcode: data.barcode,
      catalog_number: data.catalog_number,
      classification: data.classification,
      currency_code: data.currency_code,
      full_image_src: data.full_image_src,
      small_image_src: data.small_image_src,
      media_format: data.media_format,
      price: +data.price,
      publisher: data.publisher
    }

    const fullAlbum: FullAlbum = {
      credits: credits,
      tracklist: tracklist,
      album:album,
      info:info
    }

    create(fullAlbum)
    // console.log(JSON.stringify(fullAlbum, null, 2));
  }

  //TODO:create for each
  let handleChangeTitle = (i: number, e: ChangeEvent<HTMLInputElement>) => {
    let newFormValues = [...tracklist]
    newFormValues[i].title = e.target.value
    setTracklist(newFormValues)
  }

  let handleChangeDuration = (i: number, e: ChangeEvent<HTMLInputElement>) => {
    let newFormValues = [...tracklist]
    newFormValues[i].duration = e.target.value
    setTracklist(newFormValues)
  }

  let handleChangeName = (i: number, e: ChangeEvent<HTMLInputElement>) => {
    let newFormValues = [...credits]
    newFormValues[i].profession = e.target.value
    setCredits(newFormValues)
  }

  let handleChangePosition = (i: number, e: ChangeEvent<HTMLInputElement>) => {
    let newFormValues = [...credits]
    newFormValues[i].person_id = +e.target.value
    setCredits(newFormValues)
  }

  let addTracklistFields = () => {
    setTracklist([...tracklist, { title: "", duration: "" }])
  }

  let addCreditFields = () => {
    setCredits([...credits, { profession: "", person_id: -1 }])
  }

  let removeFormTracklist = (i: number) => {
    let newFormValues = [...tracklist];
    newFormValues.splice(i, 1);
    setTracklist(newFormValues)
  }

  let removeFormCredits = (i: number) => {
    let newFormValues = [...credits];
    newFormValues.splice(i, 1);
    setCredits(newFormValues)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <h1>Create Album</h1>
      <label>Title</label>
      <input type="text" {...register("title")} />
      <label>Released at</label>
      <div>
        <DateInput
          value={date}
          onChange={setDate}
        />
        <br />
        <span>{new Number(date).toString()}</span>
      </div>
      <label>Catalog Number</label>
      <input type="text" {...register("catalog_number")} />
      <label>Small Image</label>
      <input type = "text" {...register("small_image_src")} />
      <label>Full Image</label>
      <input type = "text" {...register("full_image_src")} />
      <label>Barcode</label>
      <input type = "text" {...register("barcode")} />
      <label>Price</label>
      <input type='number' step="0.1" {...register("price")} />
      <label>Currency code</label>
      <input type = "text" {...register("currency_code")} />
      <label>Media format</label>
      <input type = "text" {...register("media_format")} />
      <label>Classification</label>
      <input type = "text" {...register("classification")} />
      <label>Publisher</label>
      <input type = "text" {...register("publisher")} />
      {tracklist.map((element, index) => (
        <div className="form-inline" key={index}>
          <label>Title</label>
          <input type="text" name="title" value={element.title || ""} onChange={e => handleChangeTitle(index, e)} />
          <label>Duration</label>
          <input type="text" name="duration" value={element.duration || ""} onChange={e => handleChangeDuration(index, e)} />
          {
            index ?
              <button type="button" className="button remove" onClick={() => removeFormTracklist(index)}>Remove</button>
              : null
          }
        </div>
      ))}
      <div className="button-section">
        <button className="button add" type="button" onClick={() => addTracklistFields()}>Add</button>
        <button className="button submit" type="submit">Submit</button>
      </div>
      {credits.map((element, index) => (
        <div className="form-inline" key={index}>
          <label>Name</label>
          <input type="text" name="name" value={element.profession || ""} onChange={e => handleChangeName(index, e)} />
          <label>Position</label>
          <input type="text" name="position" value={element.person_id || 0} onChange={e => handleChangePosition(index, e)} />
          {
            index ?
              <button type="button" className="button remove" onClick={() => removeFormCredits(index)}>Remove</button>
              : null
          }
        </div>
      ))}
      <div className="button-section">
        <button className="button add" type="button" onClick={() => addCreditFields()}>Add</button>
        <button className="button submit" type="submit">Submit</button>
      </div>
      <input type="submit" />
    </form>
  );
}

export default CreateForm