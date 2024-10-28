import { FaCheck } from "react-icons/fa6";

function ButtonRound({ onClick }) {
  return (
    <div>
      <button onClick={() => onClick()} className="bg-yellow w-14 h-14 flex items-center justify-center p-4 rounded-[50%] shadow-lg hover:bg-blue"><FaCheck /></button>
    </div>
  );
}

export default ButtonRound;
