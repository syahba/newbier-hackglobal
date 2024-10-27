function FieldActivity() {
  return (
    <div>
      <h5 className="mb-4 text-sm font-bold">
        Let’s start with your preferred activities!
      </h5>
      
      <div className="flex flex-wrap gap-2">
        <div>
          <input
            type="radio"
            name="activity"
            id="activity-sightseeing"
            className="peer hidden"
          />
          <label
            htmlFor="activity-sightseeing"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Sightseeing
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            id="activity-culinary-experiences"
            className="peer hidden"
          />
          <label
            htmlFor="activity-culinary-experiences"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Culinary Experiences
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            id="activity-outdoor-activities"
            className="peer hidden"
          />
          <label
            htmlFor="activity-outdoor-activities"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Outdoor Activities
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            id="activity-adventure"
            className="peer hidden"
          />
          <label
            htmlFor="activity-adventure"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Adventure
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            id="activity-cultural-experiences"
            className="peer hidden"
          />
          <label
            htmlFor="activity-cultural-experiences"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Cultural Experiences
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            id="activity-shopping"
            className="peer hidden"
          />
          <label
            htmlFor="activity-shopping"
            className="border-blue peer-checked:bg-blue inline-flex rounded-lg border px-4 py-3 peer-checked:text-white"
          >
            Shopping
          </label>
        </div>
      </div>
    </div>
  );
}

export default FieldActivity;
