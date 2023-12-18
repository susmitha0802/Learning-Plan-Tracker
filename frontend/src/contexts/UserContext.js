import React, { createContext, useContext, useEffect, useState } from "react";

const UserContext = createContext();

export const useUser = () => {
    return useContext(UserContext);
};

export const UserProvider = ({ children }) => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        const storedUsers = JSON.parse(localStorage.getItem("users")) || [];
        setUsers(storedUsers);
    }, []);

    const addUserToLocalStorage = (email) => {
        const newUser = { email, submitted: [] };
        const updatedUsers = [...users, newUser];
        setUsers(updatedUsers);
        localStorage.setItem("users", JSON.stringify(updatedUsers));
    };

    const setsubmittedInLocalStorage = (userEmail, courseId, topicId, exerciseId, fileName, data) => {
        const localStorageKey = `exercise-${courseId}-${topicId}-${exerciseId}`;
        const newItem = { [localStorageKey]: data, courseId, topicId, exerciseId, fileName };
        const updatedUsers = users.map((user) => {
            if (user.email === userEmail) {
                if (!user.submitted.some((item) => item.localStorageKey === localStorageKey)) {
                    user.submitted.push(newItem);
                }
            }
            return user;
        });
        setUsers(updatedUsers);
        localStorage.setItem("users", JSON.stringify(updatedUsers));
    };

    const getsubmittedByUserEmail = (userEmail) => {
        const user = users.find((user) => user.email === userEmail);
        return user ? user.submitted : [];
    };

    const getSubmittedFromLocalStorage = (userEmail, courseId, topicId, exerciseId) => {
        const submittedExercises = getsubmittedByUserEmail(userEmail);
        const targetKey = `exercise-${courseId}-${topicId}-${exerciseId}`;
        const storedData = submittedExercises.find((submitted) => targetKey in submitted)
        return storedData;
    };

    const getSubmittedCountFromLocalStorage = (userEmail, id) => {
        const submittedExercises = getsubmittedByUserEmail(userEmail);
        var count = 0;
        submittedExercises.forEach(submit => {
            if (submit.courseId === JSON.parse(id)) {
                count = count + 1;
            }
        });
        return count;
    };

    const removeFromsubmittedInLocalStorage = (userEmail, courseId, topicId, exerciseId) => {
        const submittedExercises = getsubmittedByUserEmail(userEmail);
        const storedData = getSubmittedFromLocalStorage(userEmail, courseId, topicId, exerciseId);
        const updatedUsers = users.map((user) => {
            if (user.email === userEmail && user.submitted && storedData) {
                const index = submittedExercises.indexOf(storedData)
                user.submitted.splice(index, 1);
            }
            return user;
        });
        setUsers(updatedUsers);
        localStorage.setItem("users", JSON.stringify(updatedUsers));
    };

    const value = {
        users,
        addUserToLocalStorage,
        setsubmittedInLocalStorage,
        getsubmittedByUserEmail,
        getSubmittedFromLocalStorage,
        getSubmittedCountFromLocalStorage,
        removeFromsubmittedInLocalStorage
    };

    return (
        <UserContext.Provider value={value}>
            {children}
        </UserContext.Provider>
    );
};