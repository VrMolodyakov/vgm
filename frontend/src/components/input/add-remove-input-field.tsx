import { ChangeEvent, useState } from "react"

type Element = {
    title: "", email : ""
}

function AddRemoveInputField() {

    const [formValues, setFormValues] = useState([{ title: "", duration : ""}])

    let handleChange = (i:number, e:ChangeEvent<HTMLInputElement>) => {
        let newFormValues = [...formValues];
        let n:string = e.target.name
        newFormValues[i].title = e.target.value;
        setFormValues(newFormValues);
      }
    
    let addFormFields = () => {
        setFormValues([...formValues, { title: "", duration: "" }])
      }
    
    let removeFormFields = (i:number) => {
        let newFormValues = [...formValues];
        newFormValues.splice(i, 1);
        setFormValues(newFormValues)
    }
    
    let handleSubmit = (event:any) => {
        event.preventDefault();
        alert(JSON.stringify(formValues));
    }

    return (
        <form  onSubmit={handleSubmit}>
          {formValues.map((element, index) => (
            <div className="form-inline" key={index}>
              <label>Name</label>
              <input type="text" name="name" value={element.title || ""} onChange={e => handleChange(index, e)} />
              <label>Email</label>
              <input type="text" name="email" value={element.duration || ""} onChange={e => handleChange(index, e)} />
              {
                index ? 
                  <button type="button"  className="button remove" onClick={() => removeFormFields(index)}>Remove</button> 
                : null
              }
            </div>
          ))}
          <div className="button-section">
              <button className="button add" type="button" onClick={() => addFormFields()}>Add</button>
              <button className="button submit" type="submit">Submit</button>
          </div>
      </form>
    )
}
export default AddRemoveInputField