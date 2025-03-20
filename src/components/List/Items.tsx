import { VersionData } from '@/utils/parser/types';
import { setSearchParams } from '@/utils/searchParams';

interface ItemProps {
    name: string;
    data: VersionData[string];
    index: number;
    setInfo: (data: string) => void;
    setModalOpen: (open: boolean) => void;
}

const Item: React.FC<ItemProps> = ({ name, data, setInfo, setModalOpen }) => {
    const remain = data.versions.length - 2;
    const openModal = () => {
        setInfo(name);
        setSearchParams('pkg', name);
        setModalOpen(true);
    };
    return (
        <div className="ba flex min-h-32 flex-row items-center gap-4 overflow-clip rounded-xl border border-gray-300 bg-gray-50/60 p-4 shadow-md transition-all duration-300 hover:bg-white hover:shadow-xl">
            <div className="w-1/3">
                <button
                    className="btn-a cursor-pointer text-2xl leading-9 font-bold text-wrap text-gray-900"
                    onClick={openModal}
                >
                    {name}
                </button>
            </div>
            <div className="w-2/3 text-left">
                {data.versions
                    .filter((_, index) => {
                        return index < 2;
                    })
                    .map((ver, index) => {
                        return (
                            <div
                                key={index}
                                className="flex flex-row items-center gap-4 overflow-hidden text-nowrap overflow-ellipsis whitespace-nowrap"
                            >
                                <span className="min-w-16 overflow-hidden text-left text-lg leading-9 font-bold text-nowrap overflow-ellipsis whitespace-nowrap">
                                    <span
                                        data-tooltip-id="default-tooltip"
                                        data-tooltip-content={ver.c}
                                        data-tooltip-place="top"
                                    >
                                        {ver.c}
                                    </span>
                                </span>
                                <span className="overflow-hidden leading-9 text-nowrap overflow-ellipsis whitespace-nowrap">
                                    <span
                                        data-tooltip-id="default-tooltip"
                                        data-tooltip-content={ver.go.join(
                                            ' / ',
                                        )}
                                        data-tooltip-place="top"
                                    >
                                        {ver.go.join(' / ')}
                                    </span>
                                </span>
                            </div>
                        );
                    })}
                {remain > 0 && (
                    <button
                        onClick={openModal}
                        className="btn-a cursor-pointer text-sky-500"
                    >
                        And {remain} more...
                    </button>
                )}
            </div>
        </div>
    );
};

export default Item;
