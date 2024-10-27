function FieldDate() {
  return (
    <div>
      <h5 className="mb-4 text-sm font-bold">
        What time would you like to travel?
      </h5>
      <input
        type="date"
        className="border-blue w-full rounded-lg border p-1.5 text-center"
      />
    </div>
  );
}

export default FieldDate;
