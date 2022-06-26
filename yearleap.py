import datetime
import sys

def isleap(year):
    year = int(year)
    try:
        date_obj = datetime.datetime(year,2,29)
        return True
    except:
        pass
    return False

if len(sys.argv) > 1:
    year = sys.argv[1]
    print(isleap(year))

# for i in range(20):
#     print(isleap(2000-i))