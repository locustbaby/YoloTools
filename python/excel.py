import pandas as pd

# 读取excel文件
excel_file = pd.ExcelFile('path/to/file.xlsx')

# 遍历所有sheet并将数据存储在字典中
sheet_data = {}
for sheet_name in excel_file.sheet_names:
    sheet_data[sheet_name] = excel_file.parse(sheet_name)

# 打印所有sheet数据
for sheet_name, data in sheet_data.items():
    print(f"Sheet: {sheet_name}")
    print(data)
