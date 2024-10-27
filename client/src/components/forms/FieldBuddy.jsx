function FieldBuddy({setState, disabled, value}) {
  
  const action = (event) => {
    setState(event.target.value === "yes")
  }

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
            className="relative flex items-center cursor-pointer"
            htmlFor="yes"
          >
            <input
              name="option"
              type="radio"
              className="w-5 h-5 transition-all border rounded-full appearance-none cursor-pointer peer border-gray checked:border-blue"
              id="yes"
              value="yes"
              onChange={action}
              checked={"yes" === value}
              disabled={disabled}
            />
            <span className="absolute w-3 h-3 transition-opacity duration-200 transform -translate-x-1/2 -translate-y-1/2 rounded-full opacity-0 left-1/2 top-1/2 bg-blue peer-checked:opacity-100">
              {" "}
            </span>
          </label>
          <label className="ml-2 text-sm cursor-pointer" htmlFor="yes">
            Yes
          </label>
        </div>

        <div className="inline-flex items-center">
          <label
            className="relative flex items-center cursor-pointer"
            htmlFor="no"
          >
            <input
              name="option"
              type="radio"
              className="w-5 h-5 transition-all border rounded-full appearance-none cursor-pointer peer border-gray checked:border-blue"
              id="no"
              value="no"
              onChange={action}
              checked={"no" === value}
              disabled={disabled}
            />
            <span className="absolute w-3 h-3 transition-opacity duration-200 -translate-x-1/2 -translate-y-1/2 rounded-full opacity-0 left-1/2 top-1/2 transhtmlForm bg-blue peer-checked:opacity-100">
              {" "}
            </span>
          </label>
          <label className="ml-2 text-sm cursor-pointer" htmlFor="no">
            No
          </label>
        </div>
      </div>
    </div>
  );
}

export default FieldBuddy;
