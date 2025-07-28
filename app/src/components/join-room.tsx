import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";

export function JoinRoom() {
    return (
        <div className="flex flex-col gap-4">
            <Input type="text" placeholder="Paste your room link here." />
            <Button variant="ghost" className="w-full cursor-pointer">Join Room</Button>
        </div>
    )
}