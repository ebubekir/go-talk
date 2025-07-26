'use client';

import {
    createContext,
    useContext,
    useEffect,
    useState,
    ReactNode,
} from 'react';
import {
    GoogleAuthProvider,
    signInWithEmailAndPassword,
    signInWithPopup,
    signOut,
} from 'firebase/auth';
import { auth } from '@/lib/firebase';
import { usePathname, useRouter } from 'next/navigation';
import { UserResponse } from '@/api/types';
import { api } from '@/api/api';
import {LoadingSplash} from "@/components/loading-splash";

interface AuthContextType {
    user: UserResponse | null;
    loading: boolean;
    signInWithGoogle: () => Promise<void>;
    logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | null>(null);

export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<UserResponse | null>(null);
    const [loading, setLoading] = useState(true);
    const pathname = usePathname();
    const router = useRouter();

    useEffect(() => {
        // Listen for auth state changes
        const unsubscribe = auth.onAuthStateChanged((firebaseUser) => {
            if (!firebaseUser && !pathname.startsWith('/login')) {
                router.push('/login');
                // setLoading(false);
            } else if (pathname.startsWith('/login')) {
                setLoading(false);
            } else {
                api.users
                    .getUser()
                    .then((user) => {
                        setUser(user);
                        console.log('user', user);
                        setLoading(false);
                    })
                    .catch((err: Error) => {
                        auth.signOut();
                        router.push('/login');
                    });
                setLoading(false);
            }
        });

        // Cleanup subscription
        return () => unsubscribe();
    }, [pathname]);

    const signInWithGoogle = async () => {
        try {
            const provider = new GoogleAuthProvider();
            await signInWithPopup(auth, provider);
            router.push('/');
        } catch (error) {
            console.error('Google sign in error:', error);
            throw error;
        }
    };


    const logout = async () => {
        try {
            await signOut(auth);
            router.push("/login");
        } catch (error) {
            console.error('Logout error:', error);
            throw error;
        }
    };

    const value = {
        user,
        loading,
        signInWithGoogle,
        logout,
    };

    return (
        <AuthContext.Provider value={value}>
            {loading ? <LoadingSplash /> : children}
        </AuthContext.Provider>
    );
}

// Custom hook to use auth context
export function useAuth() {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
}
