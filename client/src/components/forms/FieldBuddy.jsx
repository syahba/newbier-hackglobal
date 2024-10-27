function FieldBuddy() {
  return (
    <div>
      <h5 className="mb-1 text-sm font-bold">
        Lastly, would you like a travel buddy?
      </h5>
      <p className="mb-4 text-xs">
        Travel buddy is a lyf resident that has a similar travel plan as you. By
        choosing yes, you’d post your travel plan and be notified once a match
        as been found!
      </p>

      <div className="flex justify-around">
        <div className="inline-flex items-center">
          <label
            className="relative flex cursor-pointer items-center"
            htmlFor="yes"
          >
            <input
              name="option"
              type="radio"
              className="peer h-5 w-5 cursor-pointer appearance-none rounded-full border border-gray transition-all checked:border-blue"
              id="yes"
              value="yes"
            />
            <span className="absolute left-1/2 top-1/2 h-3 w-3 -translate-x-1/2 -translate-y-1/2 transform rounded-full bg-blue opacity-0 transition-opacity duration-200 peer-checked:opacity-100">
              {" "}
            </span>
          </label>
          <label className="ml-2 cursor-pointer text-sm" htmlFor="yes">
            Yes
          </label>
        </div>

        <div className="inline-flex items-center">
          <label
            className="relative flex cursor-pointer items-center"
            htmlFor="no"
          >
            <input
              name="option"
              type="radio"
              className="peer h-5 w-5 cursor-pointer appearance-none rounded-full border border-gray transition-all checked:border-blue"
              id="no"
              value="no"
            />
            <span className="absolute left-1/2 top-1/2 h-3 w-3 -translate-x-1/2 -translate-y-1/2 transhtmlForm rounded-full bg-blue opacity-0 transition-opacity duration-200 peer-checked:opacity-100">
              {" "}
            </span>
          </label>
          <label className="ml-2 cursor-pointer text-sm" htmlFor="no">
            No
          </label>
        </div>
      </div>
    </div>
  );
}

export default FieldBuddy;
