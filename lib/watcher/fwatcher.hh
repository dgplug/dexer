#ifndef FWATCHER_HH
#define FWATCHER_HH

#include <experimental/filesystem>
#include <chrono>
#include <thread>
#include <unordered_map>
#include <string>
#include <functional>

namespace chr = std::chrono;
namespace fs = std::experimental::filesystem;

namespace dexer{
    enum class file_status{
        create, modified, erased
    };
    class fwatcher{
        public:
            fwatcher(std::string path, chr::duration<int, std::milli> delay);
            void start(const std::function<void(std::string, dexer::file_status)> &action);

        private:
            std::unordered_map<std::string, fs::file_time_type> _paths;
            bool _running = true;
            std::string search_path;
            chr::duration<int, std::milli> _delay;

            bool contains(const std::string &key);
    };
}

#endif
