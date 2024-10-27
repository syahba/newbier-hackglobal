function FieldStyle() {
  return (
    <div>
      <h5 className="mb-4 text-sm font-bold">
        Finally, what kind of trip would you like?
      </h5>
      <div className="mb-4 flex flex-wrap gap-2">
        <div>
          <input
            type="radio"
            name="trip"
            id="trip-relaxation"
            className="peer hidden"
          />
          <label
            htmlFor="trip-relaxation"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Relaxation
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            id="trip-budget"
            className="peer hidden"
          />
          <label
            htmlFor="trip-budget"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Budget
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            id="trip-luxury"
            className="peer hidden"
          />
          <label
            htmlFor="trip-luxury"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Luxury
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            id="trip-family-friendly"
            className="peer hidden"
          />
          <label
            htmlFor="trip-family-friendly"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Family Friendly
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="trip"
            id="trip-adventure"
            className="peer hidden"
          />
          <label
            htmlFor="trip-adventure"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Adventure
          </label>
        </div>
      </div>
    </div>
  );
}

export default FieldStyle;
