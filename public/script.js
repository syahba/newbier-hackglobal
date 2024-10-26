const checked = () => {
  const radio = document.querySelector('input[name="option"]:checked').value;
  const form = document.getElementById('form-buddy');
  const btn = document.getElementById('btn-option');

  if (radio === 'yes') {
    form.classList.remove('hidden');
    btn.classList.add('hidden');
  }
}