import { useState } from "react";
import { FaStar } from "react-icons/fa";
import { FaLocationDot } from "react-icons/fa6";

function CardProduct({ data }) {
  const [hidden, setHidden] = useState(true);

  function formatTitle(title) {
    const len = 20;
    if (title.length > len) {
      return title.slice(0, len) + "...";
    }
    return title;
  }

  const dropdown = () => {
    if (hidden) {
      setHidden(false);
    } else {
      setHidden(true);
    }
  };

  return (
    <div className="rounded shadow-md">
      <div className="flex flex-grow gap-3 p-3">
        <div className="w-16 h-16 rounded bg-blue"></div>
        <div className="flex flex-col flex-grow gap-1">
          <h5 className="text-sm font-bold">{formatTitle(data.name)}</h5>
          <div className="flex gap-2 mb-1">
            <FaStar className="w-4 h-4 text-yellow" />
            {data.star}
          </div>
          <div className="flex gap-2">
            <FaLocationDot className="w-4 h-4 text-yellow" />
            Singapore
          </div>
        </div>
        <div className="self-center">
          <svg
            className={`ms-3 h-2.5 w-2.5 ${!hidden ? "rotate-180" : ""}`}
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 10 6"
            onClick={dropdown}
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="m1 1 4 4 4-4"
            />
          </svg>
        </div>
      </div>

      <div
        id="dropdown"
        className={`${hidden ? "hidden" : ""} flex flex-col gap-2 p-3`}
      >
        {data.products.map((v) => (
          <div className="flex items-center gap-2 text-sm">
            <p className="flex-grow">{v.name}</p>
            <p>{v.price}</p>
            <input
              type="text"
              className="px-1 py-0 text-center border rounded w-7 border-blue"
              placeholder="0"
            />
          </div>
        ))}
      </div>
    </div>
  );
}

export default CardProduct;
