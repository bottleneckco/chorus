import React from 'react';
import PropTypes from 'prop-types';
import '../stylesheets/form.scss';

const Form = ({ fields, handleInput, submitForm, submitButtonText }) => {
  const renderFields = () => (
    fields.map((field) => (
      <div className="form--field" key={field.name}>
        <h3 className="form--label">{field.label}</h3>
        <input
          type="text"
          name={field.name}
          value={field.value}
          onChange={handleInput}
          className="textbox-line"
        />
      </div>
    ))
  );

  return (
    <form
      autoComplete="off"
      className="form"
      onSubmit={submitForm}
    >
      {renderFields()}
      <input
        type="submit"
        value={submitButtonText}
        className="button-solid"
      />
    </form>
  );
};

Form.propTypes = {
  fields: PropTypes.array.isRequired,
  handleInput: PropTypes.func.isRequired,
  submitForm: PropTypes.func.isRequired,
  submitButtonText: PropTypes.string.isRequired
};

export default Form;
