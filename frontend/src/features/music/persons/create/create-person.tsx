import { useForm } from "react-hook-form";
import { useState } from "react"
import DateInput from "../../../../components/date-input/date-input";
import { useMusicClient } from "../../client-provider/context/context";
import { MusicService } from "../../service/music";
import { Person } from "../types";
import { usePerson } from "../hooks/usePerson";
import "./create-person.css"

const CreatePersonForm: React.FC = () => {
  const { register, handleSubmit } = useForm<Person>();
  const [date, setDate] = useState(new Date())
  let client = useMusicClient()
  let musicService = new MusicService(client)
  const { mutate: create } = usePerson(musicService)

  async function onSubmit(data: Person){
    data.birth_date = new Number(date)
    console.log(data)
    create(data)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="form-container">
      <h1>Create Person</h1>
      <label>First Name</label>
      <input type="text" {...register("first_name")} />
      <label>Last Name</label>
      <input type="text" {...register("last_name")} />
      <label>Birth Date </label>
      <div>
        <DateInput
          value={date}
          onChange={setDate}
        />
        <br />
        <span>{new Number(date).toString()}</span>
      </div>
      <input type="submit" />
    </form>
  );
}

export default CreatePersonForm