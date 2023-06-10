import { useForm } from "react-hook-form";
import { useState } from "react"
import "./create-album.css"
import DateInput from "../../../../components/date-input/date-input";
import { useMusicClient } from "../../client-provider/context/context";
import { MusicService } from "../../service/music";
import { Person } from "../types";
import { usePerson } from "../hooks/usePerson";

const CreateForm: React.FC = () => {
  const { register, handleSubmit } = useForm<Person>();
  const [date, setDate] = useState(new Date())
  let client = useMusicClient()
  let musicService = new MusicService(client)
  const { mutate: create } = usePerson(musicService)

  async function onSubmit(data: Person){
    console.log(data)
    create(data)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
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
      
    </form>
  );
}

export default CreateForm