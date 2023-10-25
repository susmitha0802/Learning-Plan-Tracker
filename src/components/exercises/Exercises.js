import { Formik, Form, Field } from "formik";
import "./Exercises.css";

export const Exercises = ({ id, question, checked, onChange }) => {
    return (
        <Formik>
            <Form>
                <label className="mx-5 my-2 font">
                    <Field
                        type="checkbox"
                        checked={checked}
                        onChange={onChange}
                        name={id}
                    />
                    {` Exercise ${id} :`}
                    <div className="mx-5 my-2 px-5">
                        {question.split("\n").map((i, key) => {
                            return <div key={key}>{i}</div>;
                        })}
                    </div>
                </label>
            </Form>
        </Formik>
    )
}
