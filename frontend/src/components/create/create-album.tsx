import { useForm } from "react-hook-form";
import { ChangeEvent, useState } from "react"

const CreateForm: React.FC = () => {
  const { register, handleSubmit } = useForm();

  const [tracklist, setTracklist] = useState([{ title: "", duration: "" }])
  const [credits, setCredits] = useState([{ name: "", position: "" }])

  let handleChange = (i: number, e: ChangeEvent<HTMLInputElement>) => {
    let newFormValues = [...tracklist];
    let n: string = e.target.name
    newFormValues[i].title = e.target.value;
    setTracklist(newFormValues);
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

  // let handleSubmit = (event:any) => {
  //     event.preventDefault();
  //     alert(JSON.stringify(formValues));
  // }

  return (
    <form>
      <h1>Create Album</h1>
      <label>Title</label>
      <input name="title" {...register} />
      <label>Released at</label>
      <input name="released_at" {...register} />
      <label>Catalog Number</label>
      <input name="catalog_number" {...register} />
      <label>Image</label>
      <input name="image" {...register} />
      <label>Barcode</label>
      <input name="barcode" {...register} />
      <label>Price</label>
      <input name="price" {...register} />
      <label>Currency code</label>
      <input name="currency_code" {...register} />
      <label>Media format</label>
      <input name="media_format" {...register} />
      <label>Classification</label>
      <input name="classification" {...register} />
      <label>Publisher</label>
      <input name="publisher" {...register} />
      <input type="submit" />
      {tracklist.map((element, index) => (
        <div className="form-inline" key={index}>
          <label>Title</label>
          <input type="text" name="title" value={element.title || ""} onChange={e => handleChange(index, e)} />
          <label>Duration</label>
          <input type="text" name="duration" value={element.duration || ""} onChange={e => handleChange(index, e)} />
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
          <input type="text" name="name" value={element.name || ""} onChange={e => handleChange(index, e)} />
          <label>Position</label>
          <input type="text" name="position" value={element.position || ""} onChange={e => handleChange(index, e)} />
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
    </form>
  );
}

export default CreateForm