#include "../include/task4.h"

namespace task4
{
    void processStage(int stage)
    {
        switch (stage)
        {
        case 1:
            std::cout << "Processing stage " << stage << ": Success" << std::endl;
            break;
        case 2:
            throw std::runtime_error("Error in stage 2: Runtime error occurred");
        case 3:
            throw std::invalid_argument("Error in stage 3: Invalid argument provided");
        case 4:
            throw std::logic_error("Error in stage 4: Logic error detected");
        default:
            throw std::out_of_range("Error: Unknown stage " + std::to_string(stage));
        }
    }

    std::vector<std::exception_ptr> runPipeline()
    {
        std::vector<std::exception_ptr> exceptions;

        for (int stage = 1; stage <= 4; ++stage)
        {
            try
            {
                processStage(stage);
            }
            catch (...)
            {
                exceptions.push_back(std::current_exception());
            }
        }

        return exceptions;
    }

    void Do()
    {
        std::vector<std::exception_ptr> exceptions = runPipeline();

        if (exceptions.empty())
        {
            std::cout << "Pipeline completed successfully" << std::endl;
        }
        else
        {
            for (const auto &eptr : exceptions)
            {
                try
                {
                    std::rethrow_exception(eptr);
                }
                catch (const std::runtime_error &e)
                {
                    std::cerr << "Caught runtime_error: " << e.what() << std::endl;
                    // throw; // Повторно выбрасываем исключение
                }
                catch (const std::invalid_argument &e)
                {
                    std::cerr << "Caught invalid_argument: " << e.what() << std::endl;
                    // throw;
                }
                catch (const std::logic_error &e)
                {
                    std::cerr << "Caught logic_error: " << e.what() << std::endl;
                    // throw;
                }
                catch (const std::exception &e)
                {
                    std::cerr << "Caught standard exception: " << e.what() << std::endl;
                    // throw;
                }
                catch (...)
                {
                    std::cerr << "Caught unknown exception" << std::endl;
                    // throw;
                }
            }
        }
    }

}