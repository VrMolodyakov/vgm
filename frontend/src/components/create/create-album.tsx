import { useForm } from "react-hook-form";
import { ChangeEvent, useState } from "react"
import "./create-album.css"
import DateInput from "../date-input/date-input";

type MusicSubmitForm = {
  title: string
  catalogNumber: string
  image: string
  barcode: string
  price: number
  currencyCode: string
  mediaFormat: string
  classification: string
  publisher: string
}

async function onSubmit(data: MusicSubmitForm){
  console.log(JSON.stringify(data, null, 2));
};

const CreateForm: React.FC = () => {
  const { register, handleSubmit } = useForm<MusicSubmitForm>();
  const [date, setDate] = useState(new Date())
  const [tracklist, setTracklist] = useState([{ title: "", duration: "" }])
  const [credits, setCredits] = useState([{ name: "", position: "" }])

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
    newFormValues[i].name = e.target.value
    setCredits(newFormValues)
  }

  let handleChangePosition = (i: number, e: ChangeEvent<HTMLInputElement>) => {
    let newFormValues = [...credits]
    newFormValues[i].position = e.target.value
    setCredits(newFormValues)
  }

  let addTracklistFields = () => {
    setTracklist([...tracklist, { title: "", duration: "" }])
  }

  let addCreditFields = () => {
    setCredits([...credits, { name: "", position: "" }])
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
        <span>{date && date.toISOString()}</span>
      </div>
      <label>Catalog Number</label>
      <input type="text" {...register("catalogNumber")} />
      <label>Image</label>
      <input type = "text" {...register("image")} />
      <label>Barcode</label>
      <input type = "text" {...register("barcode")} />
      <label>Price</label>
      <input type = "number" {...register("price")} />
      <label>Currency code</label>
      <input type = "text" {...register("currencyCode")} />
      <label>Media format</label>
      <input type = "text" {...register("mediaFormat")} />
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
          <input type="text" name="name" value={element.name || ""} onChange={e => handleChangeName(index, e)} />
          <label>Position</label>
          <input type="text" name="position" value={element.position || ""} onChange={e => handleChangePosition(index, e)} />
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