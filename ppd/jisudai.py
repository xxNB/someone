import pandas as pd 

df = pd.read_excel('/Users/zhangxin/Desktop/ppdai/2drools/2drools/逾期映射表_极速贷_v20191011.xlsx')



class Data:
    df=df

# print(df)
# 初始化columns
def reanme_df():
    dic={}
    for i in df.columns:
        print(i)
        dic[i] = str(int(i.split(':')[1].strip(' '))-1)
    Data.df=df.rename(columns=dic)
    # print(Data.df)

def write_mysql(df):
     for x in df.index:
        for iy, y in enumerate(df.columns):
            print('===========', x+1, iy+1, df.loc[x,y])

def all_due():
    df = Data.df.loc[3:22, ['1','2', '3',  '4' ,'5']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)



def due_3():
    df = Data.df.loc[27:46, ['1','2', '3',  '4' ,'5']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_6():   
    df = Data.df.loc[27:47, [ '9','10','11','12','13']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

# def due_9():   
#     df = Data.df.loc[43:53, ['0','1','2', '3',  '4' , '5','6','7', '8','9','10']].reset_index()
#     del df['index']
#     print(df)
#     write_mysql(df)

def due_12():   
    df = Data.df.loc[27:47, ['17','18','19', '20','21']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

if __name__ == "__main__":
    reanme_df()
    # all_due()
    due_6()
    # reanme_df()