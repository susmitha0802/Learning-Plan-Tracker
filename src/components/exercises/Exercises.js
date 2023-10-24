import { Formik, Form, Field } from "formik";

export const Exercises = ({ id, question, status }) => {
    const initialValues = { status: status }

    return (
        <Formik
            initialValues={initialValues}
        >
            <Form>
                <label>
                    <Field
                        type="checkbox"
                        name="status" />
                    {`Exercise ${id} :`}
                    {question.split("\n").map((i, key) => {
                        return <div key={key}>{i}</div>;
                    })}
                </label>
            </Form>
        </Formik>
    )
}
