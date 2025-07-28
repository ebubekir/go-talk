import {Separator} from "@/components/ui/separator";

export function HorizontalOrSeparator() {
    return (
        <div className="flex w-full items-center justify-center gap-4">
            <div className="w-1/2">
                <Separator orientation="horizontal"  />

            </div>
            <p>or</p>
            <div className="w-1/2">
                <Separator orientation="horizontal"  />
            </div>
        </div>
    )
}