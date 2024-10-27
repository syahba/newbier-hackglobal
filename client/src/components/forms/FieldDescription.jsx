function FieldDescription() {
  return (
    <div>
      <h5 className="mb-1 text-sm font-bold">Tell us a bit about yourself!</h5>
      <p className="mb-4 text-xs">
        Writing about your planned activities, such as event you’d like to
        attend or rides you’d like to try, would help us get you a better match.
      </p>
      <textarea
        name="desc"
        id="desc"
        placeholder="Write your interests"
        className="h-32 w-full rounded-lg p-3 text-sm shadow-lg outline-blue"
      ></textarea>
    </div>
  );
}

export default FieldDescription;
