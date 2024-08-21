import { User } from "../types/User"

const API_URL = import.meta.env.VITE_WL_API_URL

export async function CreateUser(user: User): Promise<User> {
    try {
        const response = await fetch(API_URL + "/users", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(user),
        });
        if (!response.ok) {
            throw new Error(`Error creating user ${user.Name}`);
        }
        const returnedUser: User = await response.json();
        return returnedUser;
    } catch (error) {
        return {} as User;
    }
}

export async function GetUsers(): Promise<User[]> {
    console.log(API_URL+"/users")
    try {
        const response = await fetch(API_URL + "/users", {
            method: "GET",
        });
        if (!response.ok) {
            throw new Error("Error retrieving users");
        }
        const returnedUsers: User[] = await response.json();
        return returnedUsers;
    } catch (error) {
        return {} as User[];
    }
}

export async function UpdateUser(user: User): Promise<boolean> {
    try {
        const response = await fetch(API_URL + "/users", {
            method: "PATCH",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(user),
        });
        if (!response.ok) {
            throw new Error(`Error updating user ${user.Id}`);
        }
        return true;
    } catch (error) {
        return false;
    }
}

export async function DeleteUser(userId: string): Promise<boolean> {
    try {
        const response = await fetch(API_URL + "/users/" + userId, {
            method: "DELETE",
            credentials: "include",
        });
        if (!response.ok) {
            throw new Error(`Error deleting user ${userId}`);
        }
        return true;
    } catch (error) {
        return false;
    }
}
