function FieldStyle({parentStateSetter}) {
  const handleActivityChange = (event) => {
    parentStateSetter(event.target.value);
  };
  
  return (
    <div>
      <h5 className="mb-4 text-sm font-bold">
        Finally, what kind of trip would you like?
      </h5>
      <div className="flex flex-wrap gap-2 mb-4">
        <div>
          <input
            type="radio"
            name="trip"
            value="relaxation"
            id="trip-relaxation"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="trip-relaxation"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Relaxation
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            value="budget"
            id="trip-budget"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="trip-budget"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Budget
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            value="luxury"
            id="trip-luxury"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="trip-luxury"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Luxury
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            value="family_friendly"
            id="trip-family-friendly"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="trip-family-friendly"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Family Friendly
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            value="adventure"
            id="trip-adventure"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="trip-adventure"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Adventure
          </label>
        </div>
      </div>
    </div>
  );
}

export default FieldStyle;
