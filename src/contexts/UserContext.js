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

    const setsubmittedInLocalStorage = (
        userEmail,
        courseId,
        topicId,
        exerciseId,
        data
    ) => {
        const localStorageKey = `exercise-${courseId}-${topicId}-${exerciseId}`;
        const newItem = { [localStorageKey]: data, };
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
    const removeFromsubmittedInLocalStorage = (userEmail, itemId) => {
        const updatedUsers = users.map((user) => {
            if (user.email === userEmail) {
                // Remove item from submitted based on itemId
                user.submitted = user.submitted.filter((item) => item.id !== itemId);
            }
            return user;
        });
        setUsers(updatedUsers);
        // Update local storage with the updated users data
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

    const checkAndAddUser = (email) => {
        const userExists = users.some((user) => user.email === email);
        if (!userExists) {
            addUserToLocalStorage(email);
        }
    };
    const value = {
        users,
        checkAndAddUser,
        addUserToLocalStorage,
        setsubmittedInLocalStorage,
        getsubmittedByUserEmail,
        removeFromsubmittedInLocalStorage,
        getSubmittedFromLocalStorage
    };
    return (
        <UserContext.Provider value={value}>
            {children}
        </UserContext.Provider>
    );
};