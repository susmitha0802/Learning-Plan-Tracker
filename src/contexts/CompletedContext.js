import { createContext, useContext, useState } from "react";

const CompletedContext = createContext();

export const useCompleted = () => {
    return useContext(CompletedContext);
}

export const CompletedProvider = ({ children }) => {

    // localStorage.setItem(completedExercises, [])
    const [count, setCount] = useState(0);
    const [completedExercises, setCompletedExercises] = useState([]);

    const handleCheckboxChange = (exerciseId) => {
        const index = completedExercises.indexOf(exerciseId);
        if (index === -1) {
            setCompletedExercises([...completedExercises, exerciseId]);
            setCount(completedExercises.length + 1);
        }
        else {
            const newCompletedExercises = [...completedExercises];
            newCompletedExercises.splice(index, 1);
            setCompletedExercises(newCompletedExercises);
            setCount(completedExercises.length - 1);
        }
    }

    const value = {
        count,
        completedExercises,
        handleCheckboxChange,
    }

    return (
        <CompletedContext.Provider value={value}>
            {children}
        </CompletedContext.Provider>
    );
}