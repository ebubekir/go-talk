import { Bubbles } from 'lucide-react';


export function LoadingSplash() {
    return <div className="flex h-screen w-screen items-center justify-center bg-background">
        <Bubbles className="size-12 animate-spin" />
    </div>
}