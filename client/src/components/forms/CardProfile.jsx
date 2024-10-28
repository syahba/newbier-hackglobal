import { IoMdFemale } from "react-icons/io";
import { IoMdMale } from "react-icons/io";
import { BiWorld } from "react-icons/bi";
import { FaUser } from "react-icons/fa6";

function CardProfile({description, name, country, gender}) {
  // dummy, akan dimasukan jadi props variablenya
  const user = {
    name: name || "syahba",
    country: country || "indonesia",
    gender: gender || "female",
    description: description
  };

  return (
    <div className="flex flex-col h-auto gap-3 p-4 mt-5 mb-10 text-sm bg-white rounded-lg shadow-lg w-80">
      <h5 className="font-bold">Come join my trip!</h5>
      <div className="flex gap-4">
        <div className="flex items-center gap-2">
          <FaUser className="w-5 h-5 text-black" />
          <p>{user.name}</p>
        </div>
        <div className="flex items-center gap-2">
          <BiWorld className="w-5 h-5 text-black" />
          <p>{user.country}</p>
        </div>
        <div className="flex items-center gap-2">
          {user.gender === "female" ? <IoMdFemale /> : <IoMdMale />}
          <p>22</p>
        </div>
      </div>
      <p>{user.description}</p>
    </div>
  );
}

export default CardProfile;
