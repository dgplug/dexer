#include "fwatcher.hh"

dexer::fwatcher::fwatcher(std::string path, chr::duration<int, std::milli> delay):
    search_path(path), _delay(delay)
{
    for(auto &file : fs::recursive_directory_iterator(path)){
        _paths.at(file.path()) = fs::last_write_time(file);
    }
}

void 
dexer::fwatcher::start(const std::function<void(std::string, dexer::file_status)> &action)
{
    std::this_thread::sleep_for(_delay);

    for(auto file : _paths){
        if(!fs::exists(file.first)){
            action(file.first, dexer::file_status::erased);
            _paths.erase(file.first);
        }
    }

    for(auto &file : fs::recursive_directory_iterator(search_path)){
        auto last_write = fs::last_write_time(file);
        
        if(!contains(file.path())){
            _paths.at(file.path()) = last_write;
            action(file.path(), dexer::file_status::create);
            continue;
        }
        
        if(_paths.at(file.path()) != last_write){
            _paths.at(file.path()) = last_write;
            action(file.path(), dexer::file_status::modified);
        }
    }
}

bool 
dexer::fwatcher::contains(const std::string &key)
{
    auto iterator = _paths.find(key);
    return iterator != _paths.end();
}
