import { IoMdFemale } from "react-icons/io";
import { IoMdMale } from "react-icons/io";
import { BiWorld } from "react-icons/bi";
import { FaUser } from "react-icons/fa6";

function CardProfile() {
  // dummy, akan dimasukan jadi props variablenya
  const user = {
    name: "syahba",
    country: "indonesia",
    gender: "female",
    description: 'I enjoy extreme rides and horror house! I’m planning to try every single rides at USS and be spooked at the legendary horror house~'
  };

  return (
    <div className="flex h-auto w-80 flex-col gap-3 rounded-lg bg-white p-4 text-sm shadow-lg mb-10 mt-5">
      <h5 className="font-bold">Come join my trip to Universal Studio!</h5>
      <div className="flex gap-12">
        <div className="flex gap-2 items-center">
          <FaUser />
          <p>{user.name}</p>
        </div>
        <div className="flex gap-2 items-center">
          <BiWorld />
          <p>{user.country}</p>
        </div>
        <div className="flex gap-2 items-center">
          {user.gender === "female" ? <IoMdFemale /> : <IoMdMale />}
          <p>22</p>
        </div>
      </div>
      <p>{user.description}</p>
    </div>
  );
}

export default CardProfile;
